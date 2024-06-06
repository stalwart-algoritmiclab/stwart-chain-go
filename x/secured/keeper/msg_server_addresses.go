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
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
)

func (m msgServer) CreateAddresses(
	goCtx context.Context,
	msg *types.MsgCreateAddresses,
) (*types.MsgCreateAddressesResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.checkAddress(ctx, msg.Creator); err != nil {
		return nil, err
	}

	for _, address := range msg.Address {
		if address == "" {
			return nil, errorsmod.Wrapf(sdkerrors.ErrConflict, "address is empty")
		}

		if _, found := m.GetAddressesByAddress(ctx, address); found {
			return nil, errorsmod.Wrapf(sdkerrors.ErrConflict, "address %s already exists", msg.Address)
		}
	}

	id := m.AppendAddresses(ctx, types.Addresses{
		Creator: msg.Creator,
		Address: msg.Address,
	})

	return &types.MsgCreateAddressesResponse{
		Id: id,
	}, nil
}

func (m msgServer) UpdateAddresses(
	goCtx context.Context,
	msg *types.MsgUpdateAddresses,
) (*types.MsgUpdateAddressesResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.checkAddress(ctx, msg.Creator); err != nil {
		return nil, err
	}

	// Checks that the element exists
	val, found := m.GetAddresses(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	m.SetAddresses(ctx, types.Addresses{
		Creator: msg.Creator,
		Id:      msg.Id,
		Address: msg.Address,
	})

	return &types.MsgUpdateAddressesResponse{}, nil
}

func (m msgServer) DeleteAddresses(
	goCtx context.Context,
	msg *types.MsgDeleteAddresses,
) (*types.MsgDeleteAddressesResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := m.checkAddress(ctx, msg.Creator); err != nil {
		return nil, err
	}

	// Checks that the element exists
	val, found := m.GetAddresses(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	m.RemoveByID(ctx, msg.Id)

	return &types.MsgDeleteAddressesResponse{}, nil
}

func (m msgServer) checkAddress(ctx context.Context, address string) error {
	storeAdapter := runtime.KVStoreAdapter(m.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressesKey))
	addressesStore := prefix.NewStore(store, types.KeyPrefix(types.AddressesKey))

	iterator := storetypes.KVStorePrefixIterator(addressesStore, []byte{})
	defer iterator.Close()

	var list []string
	for ; iterator.Valid(); iterator.Next() {
		var addresses types.Addresses
		m.cdc.MustUnmarshal(iterator.Value(), &addresses)

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
