/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"

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

	var addressess []types.Addresses

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	addressesStore := prefix.NewStore(store, types.KeyPrefix(types.AddressesKey))

	pageRes, err := query.Paginate(addressesStore, req.Pagination, func(key []byte, value []byte) error {
		var addresses types.Addresses
		if err := k.cdc.Unmarshal(value, &addresses); err != nil {
			return err
		}

		addressess = append(addressess, addresses)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAddressesResponse{Addresses: addressess, Pagination: pageRes}, nil
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
