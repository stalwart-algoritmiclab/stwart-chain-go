package keeper

import (
	"context"
	"time"

	sdkioerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
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

	b := store.Get(types.StatsKey(date))
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

// AddCountUserStats add reward coins to today stats.
func (k Keeper) AddCountUserStats(ctx context.Context, countUsers uint64) {
	today := time.Now().Format(time.DateOnly)

	stats, found := k.GetStats(ctx, today)
	if !found {
		k.SetStats(ctx, types.Stats{
			Date: today,
			DailyStats: &types.DailyStats{
				CountUniqueActiveUsers: countUsers,
			},
		})
		return
	}

	stats.DailyStats.CountUniqueActiveUsers += countUsers

	k.SetStats(ctx, stats)
}

// AddNewUserToStat adds a counter for new users.
func (k Keeper) AddNewUserToStat(ctx context.Context) {
	today := time.Now().Format(time.DateOnly)

	stats, found := k.GetStats(ctx, today)
	if !found {
		k.SetStats(ctx, types.Stats{
			Date: today,
			DailyStats: &types.DailyStats{
				CountNewUsers: 1,
			},
		})
		return
	}

	stats.DailyStats.CountNewUsers++

	k.SetStats(ctx, stats)
}
