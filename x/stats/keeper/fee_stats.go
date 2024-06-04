package keeper

import (
	"context"
	"strconv"
	"time"

	sdkioerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stats/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetFeeStats set a specific feeStats in the store from its index
func (k Keeper) SetFeeStats(ctx context.Context, feeStats types.FeeStats) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeStatsKeyPrefix))
	b := k.cdc.MustMarshal(&feeStats)
	store.Set(types.FeeStatsKey(
		feeStats.Date,
	), b)
}

// GetFeeStats returns a feeStats from its index
func (k Keeper) GetFeeStats(
	ctx context.Context,
	date string,

) (val types.FeeStats, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeStatsKeyPrefix))

	b := store.Get(types.FeeStatsKey(
		date,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFeeStats removes a feeStats from the store
func (k Keeper) RemoveFeeStats(
	ctx context.Context,
	date string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeStatsKeyPrefix))
	store.Delete(types.FeeStatsKey(
		date,
	))
}

// GetAllFeeStats returns all feeStats
func (k Keeper) GetAllFeeStats(ctx context.Context) (list []types.FeeStats) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeStatsKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FeeStats
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// SetStatsNoFee set a specific stats no fee in the store from its index
func (k Keeper) SetStatsNoFee(ctx sdk.Context, amount sdk.Coins) {
	var (
		storeAdapter = runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
		store        = prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeStatsKeyPrefix))
		currentDate  = time.Now().Format(time.DateOnly)
		amountNoFee  sdk.Coins
	)

	stat, isFound := k.GetOneStatsByDate(ctx, currentDate)
	if !isFound {
		statsList := k.GetAllFeeStats(ctx)

		stat = types.FeeStats{
			Index: strconv.Itoa(len(statsList) + 1),
			Date:  currentDate,
			Stats: &types.FeeDailyStats{
				AmountWithFee: sdk.Coins{},
				AmountNoFee:   sdk.Coins{},
				Fee:           sdk.Coins{},
				CountWithFee:  0,
				CountNoFee:    0,
			},
		}
	}

	amountNoFee = stat.Stats.AmountNoFee

	stat.Stats.CountNoFee += 1
	stat.Stats.AmountNoFee = amountNoFee.Add(amount...)

	b := k.cdc.MustMarshal(&stat)
	store.Set(types.FeeStatsKey(
		stat.Index,
	), b)
}

// SetStatsFee set a specific stats with fee in the store from its index
func (k Keeper) SetStatsFee(ctx sdk.Context, amountFee sdk.Coins, amountTx sdk.Coins) {
	var (
		storeAdapter  = runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
		store         = prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeStatsKeyPrefix))
		currentDate   = time.Now().Format(time.DateOnly)
		amountWithFee sdk.Coins
		fee           sdk.Coins
	)

	stat, isFound := k.GetOneStatsByDate(ctx, currentDate)
	if !isFound {
		statsList := k.GetAllFeeStats(ctx)

		stat = types.FeeStats{
			Index: strconv.Itoa(len(statsList) + 1),
			Date:  currentDate,
			Stats: &types.FeeDailyStats{
				AmountWithFee: sdk.Coins{},
				AmountNoFee:   sdk.Coins{},
				Fee:           sdk.Coins{},
				CountWithFee:  0,
				CountNoFee:    0,
			},
		}
	}

	fee = stat.Stats.Fee
	amountWithFee = stat.Stats.AmountWithFee

	stat.Stats.Fee = fee.Add(amountFee...)
	stat.Stats.AmountWithFee = amountWithFee.Add(amountTx...)
	stat.Stats.CountWithFee += 1

	b := k.cdc.MustMarshal(&stat)
	store.Set(types.FeeStatsKey(
		stat.Index,
	), b)

}

// GetOneStatsByDate returns all stats by date
func (k Keeper) GetOneStatsByDate(ctx sdk.Context, date string) (stat types.FeeStats, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeStatsKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	parseDate, err := time.Parse(time.DateOnly, date)
	if err != nil {
		return types.FeeStats{}, false
	}

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FeeStats
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		parseStatsDate, err := time.Parse(time.DateOnly, val.Date)
		if err != nil {
			continue
		}

		if parseStatsDate.Equal(parseDate) {
			return val, true
		}
	}

	return types.FeeStats{}, false
}

// GetStatsByDate returns all stats by date
func (k Keeper) GetStatsByDate(ctx sdk.Context, startDate, endDate string) (list []types.FeeStats, err error) {
	startDateTime, err := time.Parse(time.DateOnly, startDate)
	if err != nil {
		return nil, sdkioerrors.Wrapf(types.ErrInvalidDate, "parse startDate failed: %s", err.Error())
	}

	endDateTime, err := time.Parse(time.DateOnly, endDate)
	if err != nil {
		return nil, sdkioerrors.Wrapf(types.ErrInvalidDate, "parse endDate failed: %s", err.Error())
	}

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeStatsKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FeeStats
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		dateTime, err := time.Parse(time.DateOnly, val.Date)
		if err != nil {
			continue
		}

		if dateTime.Before(startDateTime) || dateTime.After(endDateTime) {
			continue
		}

		list = append(list, val)
	}

	if len(list) == 0 {
		return nil, sdkioerrors.Wrapf(types.ErrNotFound, "stats by date")
	}

	return list, nil
}

// GetStatsByIndexes returns all stats by indexes
func (k Keeper) GetStatsByIndexes(ctx sdk.Context, startIndex, endIndex string) (list []types.FeeStats, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.FeeStatsKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.FeeStats
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Index <= endIndex && val.Index >= startIndex {
			list = append(list, val)
		}
	}

	if len(list) == 0 {
		return nil, false
	}

	return list, true
}
