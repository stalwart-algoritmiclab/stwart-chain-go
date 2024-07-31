/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateRates(goCtx context.Context, msg *types.MsgCreateRates) (*types.MsgCreateRatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Check if the value already exists
	_, isFound := k.GetRates(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "denom %s already exists", msg.Denom)
	}

	var rates = types.Rates{
		Creator:  msg.Creator,
		Denom:    msg.Denom,
		Rate:     msg.Rate,
		Decimals: msg.Decimals,
	}

	k.SetRates(
		ctx,
		rates,
	)
	return &types.MsgCreateRatesResponse{}, nil
}

func (k msgServer) UpdateRates(goCtx context.Context, msg *types.MsgUpdateRates) (*types.MsgUpdateRatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Check if the value exists
	_, isFound := k.GetRates(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var rates = types.Rates{
		Creator:  msg.Creator,
		Denom:    msg.Denom,
		Rate:     msg.Rate,
		Decimals: msg.Decimals,
	}

	k.SetRates(ctx, rates)

	return &types.MsgUpdateRatesResponse{}, nil
}

func (k msgServer) DeleteRates(goCtx context.Context, msg *types.MsgDeleteRates) (*types.MsgDeleteRatesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Check if the value exists
	_, isFound := k.GetRates(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveRates(
		ctx,
		msg.Denom,
	)

	return &types.MsgDeleteRatesResponse{}, nil
}
