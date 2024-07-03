/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"
	"time"

	sdkioerrors "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
)

// SetStats set a specific stats in the store from its index
func (k Keeper) SetStats(ctx context.Context, stats types.Stats) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StatsKeyPrefix))
	b := k.cdc.MustMarshal(&stats)
	store.Set(types.StatsKey(stats.Date), b)
}

// GetStats returns a stats from its index
func (k Keeper) GetStats(ctx context.Context, date string) (val types.Stats, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StatsKeyPrefix))

	b := store.Get(types.StatsKey(
		date,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStats removes a stats from the store
func (k Keeper) RemoveStats(ctx context.Context, date string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StatsKeyPrefix))
	store.Delete(types.StatsKey(date))
}

// GetAllStats returns all stats
func (k Keeper) GetAllStats(ctx context.Context) (list []types.Stats) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StatsKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Stats
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetStatsByDate returns all stats between startDate and endDate
func (k Keeper) GetStatsByDate(ctx sdk.Context, startDate, endDate string) (result []types.Stats, err error) {
	startDateTime, err := time.Parse(time.DateOnly, startDate)
	if err != nil {
		return nil, sdkioerrors.Wrapf(types.ErrInvalidDate, "parse startDate failed: %s", err.Error())
	}

	endDateTime, err := time.Parse(time.DateOnly, endDate)
	if err != nil {
		return nil, sdkioerrors.Wrapf(types.ErrInvalidDate, "parse endDate failed: %s", err.Error())
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StatsKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	for ; iterator.Valid(); iterator.Next() {
		var val types.Stats
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		dateTime, err := time.Parse(time.DateOnly, val.Date)
		if err != nil {
			panic("invalid stored date: " + err.Error()) // panic because of invalid stored data
		}

		if dateTime.Before(startDateTime) || dateTime.After(endDateTime) {
			continue
		}

		result = append(result, val)
	}

	return result, nil
}

// AddBurnedToDailyStats add burned coins to today`s stats.
func (k Keeper) AddBurnedToDailyStats(ctx sdk.Context, coins ...sdk.Coin) {
	dailyStats := &types.DailyStats{
		BurnedCoins: coins,
		CountBurned: uint64(len(coins)),
	}

	k.addDailyStat(ctx, dailyStats)
}

// AddIssuedToDailyStats add issued coins to today`s stats.
func (k Keeper) AddIssuedToDailyStats(ctx sdk.Context, coins ...sdk.Coin) {
	dailyStats := &types.DailyStats{
		IssuedCoins: coins,
		CountIssued: uint64(len(coins)),
	}

	k.addDailyStat(ctx, dailyStats)
}

// AddWithdrawnToDailyStats add withdrawn coins to today`s stats.
func (k Keeper) AddWithdrawnToDailyStats(ctx sdk.Context, coins ...sdk.Coin) {
	dailyStats := &types.DailyStats{
		WithdrawCoins: coins,
		CountWithdraw: uint64(len(coins)),
	}

	k.addDailyStat(ctx, dailyStats)
}

func (k Keeper) addDailyStat(ctx sdk.Context, newStats *types.DailyStats) {
	today := time.Now().Format(time.DateOnly)
	stats, found := k.GetStats(ctx, today)
	if !found {
		stats = types.Stats{
			Date: today,
			DailyStats: &types.DailyStats{
				IssuedCoins:   sdk.NewCoins(),
				CountIssued:   0,
				BurnedCoins:   sdk.NewCoins(),
				CountBurned:   0,
				WithdrawCoins: sdk.NewCoins(),
				CountWithdraw: 0,
			},
		}
	}

	up := stats.DailyStats

	if len(newStats.BurnedCoins) > 0 {
		up.BurnedCoins = sdk.NewCoins(stats.DailyStats.BurnedCoins...).Add(newStats.BurnedCoins...)
		up.CountBurned += newStats.CountBurned
	}

	if len(newStats.IssuedCoins) > 0 {
		up.IssuedCoins = sdk.NewCoins(stats.DailyStats.IssuedCoins...).Add(newStats.IssuedCoins...)
		up.CountIssued += newStats.CountIssued
	}

	if len(newStats.WithdrawCoins) > 0 {
		up.WithdrawCoins = sdk.NewCoins(stats.DailyStats.WithdrawCoins...).Add(newStats.WithdrawCoins...)
		up.CountWithdraw += newStats.CountWithdraw
	}

	stats.DailyStats = up
	k.SetStats(ctx, stats)
}
