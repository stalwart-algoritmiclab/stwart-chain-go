/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetUser(
		ctx,
		msg.AccountAddress,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var user = types.User{
		AccountAddress: msg.AccountAddress,
		Referrer:       msg.Referrer,
		Referrals:      msg.Referrals,
	}

	k.SetUser(
		ctx,
		user,
	)
	return &types.MsgCreateUserResponse{}, nil
}

func (k msgServer) UpdateUser(goCtx context.Context, msg *types.MsgUpdateUser) (*types.MsgUpdateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetUser(
		ctx,
		msg.AccountAddress,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	var user = types.User{
		AccountAddress: msg.AccountAddress,
		Referrer:       msg.Referrer,
		Referrals:      msg.Referrals,
	}

	k.SetUser(ctx, user)

	return &types.MsgUpdateUserResponse{}, nil
}

func (k msgServer) DeleteUser(goCtx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	_, isFound := k.GetUser(
		ctx,
		msg.AccountAddress,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	k.RemoveUser(
		ctx,
		msg.AccountAddress,
	)

	return &types.MsgDeleteUserResponse{}, nil
}
