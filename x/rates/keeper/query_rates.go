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
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RatesAll(ctx context.Context, req *types.QueryAllRatesRequest) (*types.QueryAllRatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var ratess []types.Rates

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	ratesStore := prefix.NewStore(store, types.KeyPrefix(types.RatesKeyPrefix))

	pageRes, err := query.Paginate(ratesStore, req.Pagination, func(key []byte, value []byte) error {
		var rates types.Rates
		if err := k.cdc.Unmarshal(value, &rates); err != nil {
			return err
		}

		ratess = append(ratess, rates)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRatesResponse{Rates: ratess, Pagination: pageRes}, nil
}

func (k Keeper) Rates(ctx context.Context, req *types.QueryGetRatesRequest) (*types.QueryGetRatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetRates(
		ctx,
		req.Denom,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetRatesResponse{Rates: val}, nil
}
