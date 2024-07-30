/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

// GetOptionsCount get the total number of options
func (k Keeper) GetOptionsCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.OptionsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetOptionsCount set the total number of options
func (k Keeper) SetOptionsCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.OptionsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendOptions appends a options in the store with a new id and update the count
func (k Keeper) AppendOptions(
	ctx context.Context,
	options types.Options,
) uint64 {
	// Create the options
	count := k.GetOptionsCount(ctx)

	// Set the ID of the appended value
	options.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OptionsKey))
	appendedValue := k.cdc.MustMarshal(&options)
	store.Set(GetOptionsIDBytes(options.Id), appendedValue)

	// Update options count
	k.SetOptionsCount(ctx, count+1)

	return count
}

// SetOptions set a specific options in the store
func (k Keeper) SetOptions(ctx context.Context, options types.Options) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OptionsKey))
	b := k.cdc.MustMarshal(&options)
	store.Set(GetOptionsIDBytes(options.Id), b)
}

// GetOptions returns a options from its id
func (k Keeper) GetOptions(ctx context.Context, id uint64) (val types.Options, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OptionsKey))
	b := store.Get(GetOptionsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveOptions removes a options from the store
func (k Keeper) RemoveOptions(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OptionsKey))
	store.Delete(GetOptionsIDBytes(id))
}

// GetAllOptions returns all options
func (k Keeper) GetAllOptions(ctx context.Context) (list []types.Options) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.OptionsKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Options
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetOptionsIDBytes returns the byte representation of the ID
func GetOptionsIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.OptionsKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
