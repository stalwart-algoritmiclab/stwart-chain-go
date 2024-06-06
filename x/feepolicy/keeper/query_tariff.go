/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) TariffAll(ctx context.Context, req *types.QueryAllTariffRequest) (*types.QueryAllTariffResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var tariffs []types.Tariff

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	tariffStore := prefix.NewStore(store, types.KeyPrefix(types.TariffKeyPrefix))

	pageRes, err := query.Paginate(tariffStore, req.Pagination, func(key []byte, value []byte) error {
		var tariff types.Tariff
		if err := k.cdc.Unmarshal(value, &tariff); err != nil {
			return err
		}

		tariffs = append(tariffs, tariff)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTariffResponse{Tariff: tariffs, Pagination: pageRes}, nil
}

func (k Keeper) Tariff(ctx context.Context, req *types.QueryGetTariffRequest) (*types.QueryGetTariffResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetTariff(
		ctx,
		req.Denom,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTariffResponse{Tariff: val}, nil
}
