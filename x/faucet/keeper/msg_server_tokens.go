/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
)

func (m msgServer) CreateTokens(
	goCtx context.Context,
	msg *types.MsgCreateTokens,
) (*types.MsgCreateTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := m.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	var tokens = types.Tokens{
		Creator: msg.Creator,
		Denom:   msg.Denom,
		Amount:  msg.Amount,
	}

	id := m.AppendTokens(ctx, tokens)

	return &types.MsgCreateTokensResponse{
		Id: id,
	}, nil
}

func (m msgServer) UpdateTokens(
	goCtx context.Context,
	msg *types.MsgUpdateTokens,
) (*types.MsgUpdateTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := m.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	var tokens = types.Tokens{
		Creator: msg.Creator,
		Id:      msg.Id,
		Denom:   msg.Denom,
		Amount:  msg.Amount,
	}

	// Checks that the element exists
	val, found := m.GetTokens(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	m.SetTokens(ctx, tokens)

	return &types.MsgUpdateTokensResponse{}, nil
}

func (m msgServer) DeleteTokens(
	goCtx context.Context,
	msg *types.MsgDeleteTokens,
) (*types.MsgDeleteTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := m.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Checks that the element exists
	val, found := m.GetTokens(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	m.RemoveTokens(ctx, msg.Id)

	return &types.MsgDeleteTokensResponse{}, nil
}
