package keeper

import (
	"context"
	"fmt"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAddresses(goCtx context.Context, msg *types.MsgCreateAddresses) (*types.MsgCreateAddressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	var addresses = types.Addresses{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	for _, address := range msg.Address {
		if _, found := k.GetAddressesByAddress(ctx, address); found {
			return nil, errorsmod.Wrapf(sdkerrors.ErrConflict, "address %s already exists", msg.Address)
		}
	}

	id := k.AppendAddresses(
		ctx,
		addresses,
	)

	return &types.MsgCreateAddressesResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAddresses(goCtx context.Context, msg *types.MsgUpdateAddresses) (*types.MsgUpdateAddressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	var addresses = types.Addresses{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
	}

	// Checks that the element exists
	val, found := k.GetAddresses(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAddresses(ctx, addresses)

	return &types.MsgUpdateAddressesResponse{}, nil
}

func (k msgServer) DeleteAddresses(goCtx context.Context, msg *types.MsgDeleteAddresses) (*types.MsgDeleteAddressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.securedKeeper.GetAddressesByAddress(ctx, msg.Creator)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "address %s is not allowed", msg.Creator)
	}

	// Checks that the element exists
	val, found := k.GetAddresses(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAddresses(ctx, msg.Id)

	return &types.MsgDeleteAddressesResponse{}, nil
}
