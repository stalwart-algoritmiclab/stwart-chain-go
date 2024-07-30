/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	poll, found := k.Keeper.GetPolls(ctx, msg.PollId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, "poll not found")
	}

	_, found = k.Keeper.GetVoteByAccountAddressAndPollID(ctx, msg.Creator, msg.PollId)
	if found {
		return nil, errorsmod.Wrap(sdkerrors.ErrConflict, "you have already voted")
	}

	if poll.Status != StatusVotingPeriod.String() {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "poll is not in voting period")
	}

	creatorAccount, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	totalDelegatedTokens := sdk.NewCoin(domain.DenomStake, sdkmath.ZeroInt())
	delegations, err := k.stakingKeeper.GetAllDelegatorDelegations(ctx, creatorAccount)
	if err != nil {
		return nil, err
	}

	unbondingDelegations, err := k.stakingKeeper.GetAllUnbondingDelegations(ctx, creatorAccount)
	if err != nil {
		return nil, err
	}

	totalUnbondingTokens := sdkmath.ZeroInt()
	for _, unbonding := range unbondingDelegations {
		for _, entry := range unbonding.Entries {
			totalUnbondingTokens = totalUnbondingTokens.Add(entry.Balance)
		}
	}

	for _, delegation := range delegations {
		delegatedTokens := delegation.Shares.TruncateInt()
		totalDelegatedTokens = totalDelegatedTokens.Add(sdk.NewCoin(domain.DenomStake, delegatedTokens))
	}

	if !totalDelegatedTokens.IsZero() {
		totalDelegatedTokens = totalDelegatedTokens.Sub(sdk.NewCoin(domain.DenomStake, totalUnbondingTokens))
	}

	freeStakeAmount := sdk.NewCoins(msg.Amount...).AmountOf(domain.DenomStake)

	endTime, err := time.Parse(time.RFC3339, poll.VotingEndTime)
	if err != nil {
		return nil, err
	}

	spendableCoins := k.bankKeeper.SpendableCoins(ctx, creatorAccount)
	if spendableCoins.AmountOf(domain.DenomStake).LT(freeStakeAmount) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "insufficient funds to vote with %v", msg.Amount)
	}

	if err := k.lockTokens(ctx, creatorAccount, msg.Amount, endTime.Unix()); err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrLogic, "lock tokens failed")
	}

	option, found := k.GetPollsOptionByID(poll, msg.OptionId)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, "option not found")
	}

	tokensAmount := sdk.NewCoins(option.TokensAmount...).Add(totalDelegatedTokens.AddAmount(freeStakeAmount))
	option.TokensAmount = tokensAmount
	option.VotersCount++

	poll, updated := UpdateOption(poll, option)
	if !updated {
		return nil, errorsmod.Wrap(sdkerrors.ErrNotFound, "option not found")
	}

	k.AppendVotes(ctx, types.Votes{
		AccountAddress: msg.Creator,
		PollId:         msg.PollId,
		OptionId:       msg.OptionId,
		Amount:         tokensAmount,
	})

	k.SetPolls(ctx, poll)

	return &types.MsgVoteResponse{}, nil
}

// lockTokens helper function for locking tokens.
func (k msgServer) lockTokens(goCtx context.Context, creatorAddress sdk.AccAddress, amount sdk.Coins, endTime int64) (err error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	timeNow := ctx.BlockTime().UTC().Unix()

	creatorAccount := k.accountKeeper.GetAccount(ctx, creatorAddress)
	if creatorAccount == nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "account %s not found", creatorAddress.String())
	}

	vestingAccount, ok := creatorAccount.(*vestingtypes.PeriodicVestingAccount)
	if !ok {
		baseAcc := creatorAccount.(*authtypes.BaseAccount)

		periods := vestingtypes.Periods{
			{
				Length: endTime - timeNow,
				Amount: amount,
			},
		}

		vestingAccount, err = vestingtypes.NewPeriodicVestingAccount(baseAcc, amount, timeNow, periods)
		if err != nil {
			return err
		}

		k.accountKeeper.SetAccount(ctx, vestingAccount)

		return nil
	}

	periodLength := endTime - timeNow
	period := vestingtypes.Period{
		Length: periodLength,
		Amount: amount,
	}

	if vestingAccount.EndTime > timeNow {
		vestingAccount.OriginalVesting = vestingAccount.OriginalVesting.Add(amount...)
		vestingAccount.VestingPeriods = append(vestingAccount.VestingPeriods, period)
		vestingAccount.EndTime = vestingAccount.EndTime + periodLength
	} else {
		vestingAccount.OriginalVesting = amount
		vestingAccount.VestingPeriods = vestingtypes.Periods{period}
		vestingAccount.EndTime = timeNow + periodLength
		vestingAccount.StartTime = timeNow
	}

	k.accountKeeper.SetAccount(ctx, vestingAccount)

	return nil
}
