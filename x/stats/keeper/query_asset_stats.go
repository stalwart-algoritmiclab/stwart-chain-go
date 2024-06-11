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

func (k Keeper) AssetStats(goCtx context.Context, req *types.QueryDateRequest) (*types.QueryAssetStatsResponse, error) {
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

	var (
		ctx   = sdk.UnwrapSDKContext(goCtx)
		start = startDate.Format(time.DateOnly)
		end   = endDate.Format(time.DateOnly)
	)

	var stats = &types.AssetDailyStats{
		AmountWithFee:     sdk.NewCoins(),
		AmountNoFee:       sdk.NewCoins(),
		Fee:               sdk.NewCoins(),
		CountWithFee:      0,
		CountNoFee:        0,
		Burned:            sdk.NewCoins(),
		CountBurned:       0,
		Issued:            sdk.NewCoins(),
		CountIssued:       0,
		Withdraw:          sdk.NewCoins(),
		CountWithdraw:     0,
		SysRefReward:      sdk.NewCoins(),
		CountSysRefReward: 0,
	}

	// Module core
	coreStats, err := k.coreKeeper.GetStatsByDate(ctx, start, end)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid get stats by date from core module: %s", err.Error())
	}

	for _, stat := range coreStats {
		newStats := stat.DailyStats
		stats.Burned = sdk.NewCoins(stats.Burned...).Add(newStats.BurnedCoins...)
		stats.CountBurned += newStats.CountBurned

		stats.Issued = sdk.NewCoins(stats.Issued...).Add(newStats.IssuedCoins...)
		stats.CountIssued += newStats.CountIssued

		stats.Withdraw = sdk.NewCoins(stats.Withdraw...).Add(newStats.WithdrawCoins...)
		stats.CountWithdraw += newStats.CountWithdraw
	}

	// Module fee excluder // TODO: add stats
	feeExcluderStats, err := k.GetStatsByDate(ctx, start, end)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid get stats by date from fee excluder module: %s", err.Error())
	}

	for _, stat := range feeExcluderStats {
		newStats := stat.Stats
		stats.AmountWithFee = sdk.NewCoins(stats.AmountWithFee...).Add(newStats.AmountWithFee...)
		stats.AmountNoFee = sdk.NewCoins(stats.AmountNoFee...).Add(newStats.AmountNoFee...)
		stats.Fee = sdk.NewCoins(stats.Fee...).Add(newStats.Fee...)
		stats.CountWithFee += newStats.CountWithFee
		stats.CountNoFee += newStats.CountNoFee
	}

	// Module stake rewards // TODO: add stake rewards
	// stakeRewardsStats, err := k.stakeRewardsKeeper.GetStatsByDate(ctx, start, end)
	// if err != nil {
	//	return nil, status.Errorf(codes.Internal, "invalid get stats by date from stake rewards module %s", err.Error())
	// }
	//
	// for _, stat := range stakeRewardsStats {
	//	newStats := stat.DailyStats
	//	stats.SysReward = sdk.NewCoins(stats.SysReward...).Add(newStats.SysReward...)
	//	stats.CountSysReward += newStats.CountSysReward
	// }

	rewardsStats, err := k.rewardsKeeper.GetStatsByDate(ctx, start, end)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "invalid get stats by date from rewards module %s", err.Error())
	}

	for _, stat := range rewardsStats {
		newStats := stat.DailyStats
		stats.SysRefReward = sdk.NewCoins(stats.SysRefReward...).Add(newStats.Reward...)
		stats.CountSysRefReward += newStats.Count
	}

	return &types.QueryAssetStatsResponse{Stats: types.AssetStats{
		DailyStats: stats,
	}}, nil
}
