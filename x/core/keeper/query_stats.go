/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"

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

	var statss []types.Stats

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	statsStore := prefix.NewStore(store, types.KeyPrefix(types.StatsKeyPrefix))

	pageRes, err := query.Paginate(statsStore, req.Pagination, func(key []byte, value []byte) error {
		var stats types.Stats
		if err := k.cdc.Unmarshal(value, &stats); err != nil {
			return err
		}

		statss = append(statss, stats)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStatsResponse{Stats: statss, Pagination: pageRes}, nil
}

func (k Keeper) Stats(ctx context.Context, req *types.QueryGetStatsRequest) (*types.QueryGetStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetStats(
		ctx,
		req.Date,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStatsResponse{Stats: val}, nil
}

func (k Keeper) StatsByDate(goCtx context.Context, req *types.QueryGetStatsByDateRequest) (*types.QueryAllStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	startDate, err := time.Parse(time.DateOnly, req.StartDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid startDate")
	}

	endDate, err := time.Parse(time.DateOnly, req.EndDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid endDate")
	}

	if startDate.After(endDate) {
		return nil, status.Error(codes.InvalidArgument, "startDate can't be after endDate")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	result := make([]types.Stats, 0)

	paginateResp, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var val types.Stats
		if err = k.cdc.Unmarshal(value, &val); err != nil {
			return status.Error(codes.Internal, "invalid stored stats")
		}

		dateTime, err := time.Parse(time.DateOnly, val.Date)
		if err != nil {
			return status.Error(codes.Internal, "invalid stored field date")
		}

		if dateTime.Before(startDate) || dateTime.After(endDate) {
			return nil
		}

		result = append(result, val)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryAllStatsResponse{Stats: result, Pagination: paginateResp}, nil
}
