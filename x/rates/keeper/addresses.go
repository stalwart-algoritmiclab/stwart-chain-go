/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetAddressesCount get the total number of addresses
func (k Keeper) GetAddressesCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.AddressesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAddressesCount set the total number of addresses
func (k Keeper) SetAddressesCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.AddressesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendAddresses appends a addresses in the store with a new id and update the count
func (k Keeper) AppendAddresses(
	ctx context.Context,
	addresses types.Addresses,
) uint64 {
	// Create the addresses
	count := k.GetAddressesCount(ctx)

	// Set the ID of the appended value
	addresses.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressesKey))
	appendedValue := k.cdc.MustMarshal(&addresses)
	store.Set(GetAddressesIDBytes(addresses.Id), appendedValue)

	// Update addresses count
	k.SetAddressesCount(ctx, count+1)

	return count
}

// SetAddresses set a specific addresses in the store
func (k Keeper) SetAddresses(ctx context.Context, addresses types.Addresses) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressesKey))
	b := k.cdc.MustMarshal(&addresses)
	store.Set(GetAddressesIDBytes(addresses.Id), b)
}

// GetAddresses returns a addresses from its id
func (k Keeper) GetAddresses(ctx context.Context, id uint64) (val types.Addresses, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressesKey))
	b := store.Get(GetAddressesIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveAddresses removes a addresses from the store
func (k Keeper) RemoveAddresses(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressesKey))
	store.Delete(GetAddressesIDBytes(id))
}

// GetAllAddresses returns all addresses
func (k Keeper) GetAllAddresses(ctx context.Context) (list []types.Addresses) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressesKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Addresses
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAddressesIDBytes returns the byte representation of the ID
func GetAddressesIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.AddressesKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}

// GetAddressesByAddress returns an address from its address
func (k Keeper) GetAddressesByAddress(ctx sdk.Context, address string) (val types.Addresses, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.AddressesKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var addresses types.Addresses
		k.cdc.MustUnmarshal(iterator.Value(), &addresses)

		for _, a := range addresses.Address {
			if a == address {
				return addresses, true
			}
		}
	}

	return val, false
}
