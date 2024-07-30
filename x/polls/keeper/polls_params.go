/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

// SetPollsParams set pollsParams in the store
func (k Keeper) SetPollsParams(ctx context.Context, pollsParams types.PollsParams) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PollsParamsKey))
	b := k.cdc.MustMarshal(&pollsParams)
	store.Set([]byte{0}, b)
}

// GetPollsParams returns pollsParams
func (k Keeper) GetPollsParams(ctx context.Context) (val types.PollsParams, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PollsParamsKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePollsParams removes pollsParams from the store
func (k Keeper) RemovePollsParams(ctx context.Context) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PollsParamsKey))
	store.Delete([]byte{0})
}
