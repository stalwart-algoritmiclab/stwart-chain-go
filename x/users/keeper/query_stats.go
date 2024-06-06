/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StatsAll(ctx context.Context, req *types.QueryAllStatsRequest) (*types.QueryAllStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var statsList []types.Stats

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	statsStore := prefix.NewStore(store, types.KeyPrefix(types.StatsKeyPrefix))

	pageRes, err := query.Paginate(statsStore, req.Pagination, func(key []byte, value []byte) error {
		var stats types.Stats
		if err := k.cdc.Unmarshal(value, &stats); err != nil {
			return err
		}

		statsList = append(statsList, stats)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStatsResponse{Stats: statsList, Pagination: pageRes}, nil
}

func (k Keeper) Stats(ctx context.Context, req *types.QueryGetStatsRequest) (*types.QueryGetStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetStats(ctx, req.Date)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStatsResponse{Stats: val}, nil
}
