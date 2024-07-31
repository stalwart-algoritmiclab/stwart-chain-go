package keeper

import (
	"context"
	"time"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

const vetoOptionText = "veto"

func (k msgServer) CreatePoll(goCtx context.Context, msg *types.MsgCreatePoll) (*types.MsgCreatePollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	timeNow := ctx.BlockTime().UTC()

	moduleParams, ok := k.Keeper.GetPollsParams(ctx)
	if !ok {
		return nil, errorsmod.Wrap(sdkerrors.ErrAppConfig, "polls module params are not set")
	}

	poll := types.Polls{
		Title:           msg.Title,
		Description:     msg.Description,
		ProposerAddress: msg.Creator,
		FailureReason:   "",
		Status:          StatusPending.String(),
	}

	votingStartTime, err := time.Parse(time.RFC3339, msg.VotingStartTime)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid voting start time")
	}

	pendingTime := timeNow.Sub(votingStartTime)
	thresholdPendingTime, err := time.ParseDuration(moduleParams.MaxDaysPending)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrAppConfig, "invalid polls module params: pending time")
	}

	if pendingTime > thresholdPendingTime {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "requested pending time %s is longer than max pending time %s", pendingTime.String(), thresholdPendingTime.String())
	}

	if votingStartTime.Before(timeNow) {
		votingStartTime = timeNow
	}

	poll.VotingStartTime = votingStartTime.Format(time.RFC3339)

	votingPeriod, err := time.ParseDuration(msg.VotingPeriod)
	if err != nil {
		return nil, err
	}

	thresholdMaxVotingPeriod, err := time.ParseDuration(moduleParams.MaxDaysDuration)
	if err != nil {
		return nil, err
	}

	if votingPeriod > thresholdMaxVotingPeriod {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "requested voting period %s is longer than max voting period %s", votingPeriod.String(), thresholdMaxVotingPeriod.String())
	}

	thresholdMinVotingPeriod, err := time.ParseDuration(moduleParams.MinDaysDuration)
	if err != nil {
		return nil, err
	}

	if votingPeriod < thresholdMinVotingPeriod {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "requested voting period %s is shorter than min voting period %s", votingPeriod.String(), thresholdMinVotingPeriod.String())
	}

	poll.VotingPeriod = msg.VotingPeriod
	poll.VotingEndTime = votingStartTime.Add(votingPeriod).Format(time.RFC3339)

	if msg.MinVoteAmount <= 0 {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "min vote amount must be greater than 0")
	}

	if msg.MinAdressesCount < 0 {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "min adresses count must be greater or equal to 0")
	}

	if msg.MinVoteCoinsAmount <= 0 {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "min vote coins amount must be greater than 0")
	}

	poll.MinVoteAmount = sdk.NewCoins(sdk.NewCoin(domain.DenomStake, math.NewInt(int64(msg.MinVoteAmount))))
	poll.MinAddressesCount = msg.MinAdressesCount
	poll.MinVotedCoinsAmount = sdk.NewCoins(sdk.NewCoin(domain.DenomStake, math.NewInt(int64(msg.MinVoteCoinsAmount))))

	poll.Options, err = validateAndCreateOptions(msg.Options)
	if err != nil {
		return nil, err
	}

	proposerAccount, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if err = k.Keeper.bankKeeper.SendCoinsFromAccountToModule(ctx, proposerAccount, types.ModuleName, moduleParams.ProposerDeposit); err != nil {
		return nil, err
	}

	k.AppendPolls(ctx, poll)

	return &types.MsgCreatePollResponse{}, nil
}

// validateAndCreateOptions validates and creates options, add default veto option
func validateAndCreateOptions(options []types.Options) ([]*types.Options, error) {
	result := make([]*types.Options, 0)
	for id, option := range options {
		option.IsVeto = false
		option.IsWinner = false
		option.Id = uint64(id)

		if option.Text == "" {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "option %d: text is empty", id)
		}
		result = append(result, &option)
	}

	result = append(result, &types.Options{
		Id:       uint64(len(result)),
		IsVeto:   true,
		IsWinner: false,
		Text:     vetoOptionText,
	})

	return result, nil
}
