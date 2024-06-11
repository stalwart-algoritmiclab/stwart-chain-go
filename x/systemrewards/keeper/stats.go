/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"
	"time"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/systemrewards/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetStats set a specific stats in the store from its index
func (k Keeper) SetStats(ctx context.Context, stats types.Stats) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StatsKeyPrefix))
	b := k.cdc.MustMarshal(&stats)
	store.Set(types.StatsKey(
		stats.Date,
	), b)
}

// GetStats returns a stats from its index
func (k Keeper) GetStats(
	ctx context.Context,
	date string,

) (val types.Stats, found bool) {
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
func (k Keeper) RemoveStats(
	ctx context.Context,
	date string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StatsKeyPrefix))
	store.Delete(types.StatsKey(
		date,
	))
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

// AddStats add reward coins to today stats.
// If coins is empty, nothing will be added.
// If coins is multiple, the same coins will be merged.
// Coins must be sorted by denom.
func (k Keeper) AddStats(ctx context.Context, coins ...sdk.Coin) {
	if len(coins) == 0 {
		return
	}

	newCoins := sdk.NewCoins(coins[0])
	if len(coins) > 1 {
		newCoins = newCoins.Add(coins[1:]...) // merge the same coins for future 'Add' operation
	}

	today := time.Now().Format(time.DateOnly)
	coinsCount := uint64(len(coins))

	stats, found := k.GetStats(ctx, today)
	if !found || stats.DailyStats == nil {
		k.SetStats(ctx, types.Stats{
			Date: today,
			DailyStats: &types.DailyStats{
				Reward: newCoins,
				Count:  coinsCount,
			},
		})
		return
	}

	rewardCoins := sdk.NewCoins(stats.DailyStats.Reward...)
	stats.DailyStats.Reward = rewardCoins.Add(newCoins...)
	stats.DailyStats.Count += coinsCount

	k.SetStats(ctx, stats)
}

// GetStatsByDate returns all stats between startDate and endDate
func (k Keeper) GetStatsByDate(ctx context.Context, startDate, endDate string) (result []types.Stats, err error) {
	startDateTime, err := time.Parse(time.DateOnly, startDate)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidDate, "parse startDate failed: %s", err.Error())
	}

	endDateTime, err := time.Parse(time.DateOnly, endDate)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInvalidDate, "parse endDate failed: %s", err.Error())
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StatsKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

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
