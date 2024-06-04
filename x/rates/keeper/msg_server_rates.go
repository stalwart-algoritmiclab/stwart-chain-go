package keeper

import (
	"context"
	"strconv"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"

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
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	rateFloat, err := strconv.ParseFloat(msg.Rate, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	var rates = types.Rates{
		Creator:  msg.Creator,
		Denom:    msg.Denom,
		Rate:     rateFloat,
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
	valFound, isFound := k.GetRates(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	rateFloat, err := strconv.ParseFloat(msg.Rate, 64)
	if err != nil {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "invalid rate")
	}

	var rates = types.Rates{
		Creator:  msg.Creator,
		Denom:    msg.Denom,
		Rate:     rateFloat,
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
	valFound, isFound := k.GetRates(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveRates(
		ctx,
		msg.Denom,
	)

	return &types.MsgDeleteRatesResponse{}, nil
}
