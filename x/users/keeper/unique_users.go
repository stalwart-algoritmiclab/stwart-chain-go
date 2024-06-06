/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"
	"time"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetUniqueUsers set a specific uniqueUsers in the store from its index
func (k Keeper) SetUniqueUsers(ctx context.Context, uniqueUsers types.UniqueUsers) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniqueUsersKeyPrefix))
	b := k.cdc.MustMarshal(&uniqueUsers)
	store.Set(types.UniqueUsersKey(uniqueUsers.Date), b)
}

// GetUniqueUsers returns a uniqueUsers from its index
func (k Keeper) GetUniqueUsers(ctx context.Context, date string) (val types.UniqueUsers, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniqueUsersKeyPrefix))

	b := store.Get(types.UniqueUsersKey(date))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveUniqueUsers removes a uniqueUsers from the store
func (k Keeper) RemoveUniqueUsers(ctx context.Context, date string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniqueUsersKeyPrefix))
	store.Delete(types.UniqueUsersKey(date))
}

// GetAllUniqueUsers returns all uniqueUsers
func (k Keeper) GetAllUniqueUsers(ctx context.Context) (list []types.UniqueUsers) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.UniqueUsersKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.UniqueUsers
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// CountUsers function counts unique user addresses for a given day and updates the storage with the unique addresses.
// It also clears outdated data when the date changes.
func (k Keeper) CountUsers(ctx context.Context, userAddresses []string) uint64 {
	if len(userAddresses) == 0 {
		return 0
	}

	currentDate := time.Now().Format(time.DateOnly)

	_, uniqueAddressSlice := k.getUniqueAddressesMapAndSlice(userAddresses...)

	storedList, found := k.GetUniqueUsers(ctx, currentDate)
	if !found {
		// Remove all old unique users data
		allUniqueUsers := k.GetAllUniqueUsers(ctx)
		for _, uniqueUsers := range allUniqueUsers {
			k.RemoveUniqueUsers(ctx, uniqueUsers.Date)
		}

		k.AddCountUserStats(ctx, uint64(len(uniqueAddressSlice)))
		// If there are no unique users for today, set the unique users for today
		k.SetUniqueUsers(ctx, types.UniqueUsers{
			Date: currentDate,
			UniqueUserAddresses: &types.UniqueUserAddresses{
				Addresses: uniqueAddressSlice,
			},
		})

		return uint64(len(uniqueAddressSlice))
	}

	mapStoreUniqueUserAddresses, _ := k.getUniqueAddressesMapAndSlice(storedList.UniqueUserAddresses.Addresses...)

	// Add new unique users to the store and update the stats for the day
	var countUniqueUsers uint64
	for _, address := range uniqueAddressSlice {
		if _, ok := mapStoreUniqueUserAddresses[address]; !ok {
			countUniqueUsers++
			storedList.UniqueUserAddresses.Addresses = append(storedList.UniqueUserAddresses.Addresses, address)
		}
	}

	if countUniqueUsers == 0 {
		return 0
	}

	k.AddCountUserStats(ctx, countUniqueUsers)
	k.SetUniqueUsers(ctx, storedList)

	return countUniqueUsers
}

// getUniqueAddressesMapAndSlice returns a map and slice of unique addresses
func (k Keeper) getUniqueAddressesMapAndSlice(addresses ...string) (map[string]struct{}, []string) {
	mapAddresses := make(map[string]struct{}, len(addresses))

	for _, address := range addresses {
		mapAddresses[address] = struct{}{}
	}

	sliceAddress := make([]string, 0, len(addresses))
	for address := range mapAddresses {
		sliceAddress = append(sliceAddress, address)
	}

	return mapAddresses, sliceAddress
}
