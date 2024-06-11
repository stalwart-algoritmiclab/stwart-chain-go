/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) AddressesAll(ctx context.Context, req *types.QueryAllAddressesRequest) (*types.QueryAllAddressesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var addresss []types.Address

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	addressesStore := prefix.NewStore(store, types.KeyPrefix(types.AddressesKey))

	pageRes, err := query.Paginate(addressesStore, req.Pagination, func(key []byte, value []byte) error {
		var address types.Address
		if err := k.cdc.Unmarshal(value, &address); err != nil {
			return err
		}

		addresss = append(addresss, address)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddressesResponse{Addresses: addresss, Pagination: pageRes}, nil
}

func (k Keeper) Addresses(ctx context.Context, req *types.QueryGetAddressesRequest) (*types.QueryGetAddressesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	addresses, found := k.GetAddresses(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAddressesResponse{Addresses: addresses}, nil
}

func (k Keeper) Address(
	goCtx context.Context,
	req *types.QueryGetAddressRequest,
) (*types.QueryGetAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	address, found := k.GetAddress(ctx, req.Address)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAddressResponse{Address: address}, nil
}

func (k Keeper) AddressByID(
	goCtx context.Context,
	req *types.QueryGetAddressByIDRequest,
) (*types.QueryGetAddressResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	address, found := k.GetAddressByID(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAddressResponse{Address: address}, nil
}
