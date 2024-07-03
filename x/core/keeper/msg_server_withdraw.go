/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
)

func (m msgServer) Withdraw(goCtx context.Context, msg *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := m.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator); !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "cretor %s is not allowed", msg.Creator)
	}

	address, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "address: %s is not correct", address)
	}

	if acc := m.Keeper.accountKeeper.GetAccount(ctx, address); acc == nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "address %s is not found", msg.Address)
	}

	amount, ok := math.NewIntFromString(msg.Amount)
	if !ok {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount: %s", msg.Amount)
	}

	if msg.Denom == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "denom is empty")
	}

	coinAmount := sdk.NewCoin(msg.Denom, amount)
	coins := sdk.NewCoins(coinAmount)

	if coinAmount.IsZero() {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid amount: %s", msg.Amount)
	}

	if coinAmount.Denom == domain.DenomStake {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, "cannot withdraw stake")
	}

	if err = m.Burn(ctx, address, coinAmount); err != nil {
		return nil, err
	}

	m.AddWithdrawnToDailyStats(ctx, coins...)
	m.AddBurnedToDailyStats(ctx, coins...)

	if err = ctx.EventManager().EmitTypedEvents(msg); err != nil {
		return nil, err
	}

	return &types.MsgWithdrawResponse{}, nil
}
