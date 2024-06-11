/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stats/types"
)

func (k Keeper) UserStats(goCtx context.Context, req *types.QueryDateRequest) (*types.QueryUserStatsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var stats types.UserStats

	startDate, err := time.Parse(time.DateOnly, req.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse(time.DateOnly, req.EndDate)
	if err != nil {
		return nil, err
	}

	if startDate.After(endDate) {
		return nil, status.Error(codes.InvalidArgument, "startDate can't be after endDate")
	}

	userStats, err := k.usersKeeper.GetStatsByDate(ctx, startDate.Format(time.DateOnly), endDate.Format(time.DateOnly))
	if err != nil {
		return nil, err
	}

	for _, u := range userStats {
		stats.CountUniqueActiveUsers += u.DailyStats.CountUniqueActiveUsers
		stats.CountNewUsers += u.DailyStats.CountNewUsers
	}

	return &types.QueryUserStatsResponse{Stats: stats}, nil
}
