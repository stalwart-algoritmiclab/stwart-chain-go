/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"
)

// IncrementTotalUsers set a specific uniqueUsers in the store from its index
func (k Keeper) IncrementTotalUsers(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TotalUsersKeyPrefix))

	total := k.GetTotalUsers(ctx)
	total.Total++

	b := k.cdc.MustMarshal(&total)
	store.Set(types.UniqueUsersKey("all"), b)
}

// GetTotalUsers returns a uniqueUsers from its index
func (k Keeper) GetTotalUsers(ctx context.Context) types.TotalUsers {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TotalUsersKeyPrefix))

	var val types.TotalUsers
	b := store.Get(types.UniqueUsersKey("all"))
	if b == nil {
		return val
	}

	k.cdc.MustUnmarshal(b, &val)
	return val
}

// SetTotalUsers set total users
func (k Keeper) SetTotalUsers(ctx sdk.Context, usersCount uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TotalUsersKeyPrefix))

	total := types.TotalUsers{Total: usersCount}
	b := k.cdc.MustMarshal(&total)
	store.Set(types.UniqueUsersKey("all"), b)
}
