/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetTariff set a specific tariff in the store from its index
func (k Keeper) SetTariff(ctx context.Context, tariff types.Tariff) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TariffKeyPrefix))
	b := k.cdc.MustMarshal(&tariff)
	store.Set(types.TariffKey(
		tariff.Denom,
	), b)
}

// GetTariff returns a tariff from its index
func (k Keeper) GetTariff(
	ctx context.Context,
	denom string,

) (val types.Tariff, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TariffKeyPrefix))

	b := store.Get(types.TariffKey(
		denom,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTariff removes a tariff from the store
func (k Keeper) RemoveTariff(
	ctx context.Context,
	denom string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TariffKeyPrefix))
	store.Delete(types.TariffKey(
		denom,
	))
}

// GetAllTariff returns all tariff
func (k Keeper) GetAllTariff(ctx context.Context) (list []types.Tariff) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TariffKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tariff
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
