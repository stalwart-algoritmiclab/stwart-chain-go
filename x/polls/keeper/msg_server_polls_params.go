/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

func (k msgServer) CreatePollsParams(goCtx context.Context, msg *types.MsgCreatePollsParams) (*types.MsgCreatePollsParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator); !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Check if the value already exists
	_, isFound := k.GetPollsParams(ctx)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "already set")
	}

	var pollsParams = types.PollsParams{
		ProposerDeposit: msg.ProposerDeposit,
		MinDaysDuration: msg.MinDaysDuration,
		MaxDaysDuration: msg.MaxDaysDuration,
		MaxDaysPending:  msg.MaxDaysPending,
		BurnVeto:        msg.BurnVeto,
		Creator:         msg.Creator,
	}

	k.SetPollsParams(
		ctx,
		pollsParams,
	)
	return &types.MsgCreatePollsParamsResponse{}, nil
}

func (k msgServer) UpdatePollsParams(goCtx context.Context, msg *types.MsgUpdatePollsParams) (*types.MsgUpdatePollsParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator); !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Check if the value exists
	valFound, isFound := k.GetPollsParams(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var pollsParams = types.PollsParams{
		ProposerDeposit: msg.ProposerDeposit,
		MinDaysDuration: msg.MinDaysDuration,
		MaxDaysDuration: msg.MaxDaysDuration,
		MaxDaysPending:  msg.MaxDaysPending,
		BurnVeto:        msg.BurnVeto,
		Creator:         msg.Creator,
	}

	k.SetPollsParams(ctx, pollsParams)

	return &types.MsgUpdatePollsParamsResponse{}, nil
}

func (k msgServer) DeletePollsParams(goCtx context.Context, msg *types.MsgDeletePollsParams) (*types.MsgDeletePollsParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator); !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Check if the value exists
	valFound, isFound := k.GetPollsParams(ctx)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemovePollsParams(ctx)

	return &types.MsgDeletePollsParamsResponse{}, nil
}
