package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAddresses(goCtx context.Context, msg *types.MsgCreateAddresses) (*types.MsgCreateAddressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var addresses = types.Addresses{
		Creator: msg.Creator,
		Address: msg.Address,
	}

	if err := k.checkAddress(ctx, msg.Creator); err != nil {
		return nil, err
	}

	for _, address := range msg.Address {
		if address == "" {
			return nil, errorsmod.Wrapf(sdkerrors.ErrConflict, "address is empty")
		}

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

	var addresses = types.Addresses{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
	}

	if err := k.checkAddress(ctx, msg.Creator); err != nil {
		return nil, err
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

	if err := k.checkAddress(ctx, msg.Creator); err != nil {
		return nil, err
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

	k.RemoveByID(ctx, msg.Id)

	return &types.MsgDeleteAddressesResponse{}, nil
}

func (k msgServer) checkAddress(ctx context.Context, address string) error {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressesKey))
	addressesStore := prefix.NewStore(store, types.KeyPrefix(types.AddressesKey))

	iterator := storetypes.KVStorePrefixIterator(addressesStore, []byte{})
	defer iterator.Close()

	var list []string
	for ; iterator.Valid(); iterator.Next() {
		var addresses types.Addresses
		k.cdc.MustUnmarshal(iterator.Value(), &addresses)

		for _, a := range addresses.Address {
			list = append(list, a)
			if a == address {
				return nil
			}
		}
	}

	if len(list) == 0 {
		return nil // it is ok
	}

	return errorsmod.Wrap(sdkerrors.ErrUnauthorized, "this address is not allowed to send transactions")
}

func (k msgServer) DeleteByAddresses(
	goCtx context.Context,
	msg *types.MsgDeleteByAddresses,
) (*types.MsgDeleteAddressesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.checkAddress(ctx, msg.Creator); err != nil {
		return nil, err
	}

	// Checks that the element exists
	for _, address := range msg.Address {
		val, found := k.GetAddressesByAddress(ctx, address)
		if !found {
			return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "address %s doesn't exist", msg.Address)
		}

		// Checks if the msg creator is the same as the current owner
		if msg.Creator != val.Creator {
			return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
		}

		k.RemoveByID(ctx, val.Id)
	}

	return &types.MsgDeleteAddressesResponse{}, nil
}
