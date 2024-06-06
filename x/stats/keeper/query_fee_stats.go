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

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stats/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FeeStatsAll(ctx context.Context, req *types.QueryAllFeeStatsRequest) (*types.QueryAllFeeStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var feeStatss []types.FeeStats

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	feeStatsStore := prefix.NewStore(store, types.KeyPrefix(types.FeeStatsKeyPrefix))

	pageRes, err := query.Paginate(feeStatsStore, req.Pagination, func(key []byte, value []byte) error {
		var feeStats types.FeeStats
		if err := k.cdc.Unmarshal(value, &feeStats); err != nil {
			return err
		}

		feeStatss = append(feeStatss, feeStats)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFeeStatsResponse{FeeStats: feeStatss, Pagination: pageRes}, nil
}

func (k Keeper) FeeStats(ctx context.Context, req *types.QueryGetFeeStatsRequest) (*types.QueryGetFeeStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, found := k.GetFeeStats(
		ctx,
		req.Date,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetFeeStatsResponse{FeeStats: val}, nil
}

func (k Keeper) StatsByIndexes(
	goCtx context.Context,
	req *types.QueryGetFeeStatsByIndexesRequest,
) (*types.QueryAllFeeStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var stats []types.FeeStats
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	statsStore := prefix.NewStore(store, types.KeyPrefix(types.FeeStatsKeyPrefix))

	paginateResp, err := query.Paginate(statsStore, req.Pagination, func(key []byte, value []byte) error {
		var stat types.FeeStats
		if err := k.cdc.Unmarshal(value, &stat); err != nil {
			return err
		}

		if stat.Index <= req.EndIndex && stat.Index >= req.StartIndex {
			stats = append(stats, stat)

			return nil
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryAllFeeStatsResponse{FeeStats: stats, Pagination: paginateResp}, nil
}

func (k Keeper) StatsByDate(
	goCtx context.Context,
	req *types.QueryGetFeeStatsByDateRequest,
) (*types.QueryAllFeeStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	var stats []types.FeeStats
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	statsStore := prefix.NewStore(store, types.KeyPrefix(types.FeeStatsKeyPrefix))

	parseStartDate, err := time.Parse(time.DateOnly, req.StartDate)
	if err != nil {
		return nil, err
	}

	parseEndDate, err := time.Parse(time.DateOnly, req.EndDate)
	if err != nil {
		return nil, err
	}

	paginateResp, err := query.Paginate(statsStore, req.Pagination, func(key []byte, value []byte) error {
		var stat types.FeeStats
		if err := k.cdc.Unmarshal(value, &stat); err != nil {
			return err
		}

		parseStatsDate, err := time.Parse(time.DateOnly, stat.Date)
		if err != nil {
			return err
		}

		if parseStatsDate.Before(parseEndDate.Add(time.Hour*24)) &&
			(parseStatsDate.After(parseStartDate) || parseStatsDate.Equal(parseStatsDate)) {
			stats = append(stats, stat)

			return nil
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &types.QueryAllFeeStatsResponse{FeeStats: stats, Pagination: paginateResp}, nil
}
