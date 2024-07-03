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

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/types"
)

// GetTokensCount get the total number of tokens
func (k Keeper) GetTokensCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.TokensCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTokensCount set the total number of tokens
func (k Keeper) SetTokensCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.TokensCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTokens appends a tokens in the store with a new id and update the count
func (k Keeper) AppendTokens(
	ctx context.Context,
	tokens types.Tokens,
) uint64 {
	// Create the tokens
	count := k.GetTokensCount(ctx)

	// Set the ID of the appended value
	tokens.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokensKey))
	appendedValue := k.cdc.MustMarshal(&tokens)
	store.Set(GetTokensIDBytes(tokens.Id), appendedValue)

	// Update tokens count
	k.SetTokensCount(ctx, count+1)

	return count
}

// SetTokens set a specific tokens in the store
func (k Keeper) SetTokens(ctx context.Context, tokens types.Tokens) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokensKey))
	b := k.cdc.MustMarshal(&tokens)
	store.Set(GetTokensIDBytes(tokens.Id), b)
}

// GetTokens returns a tokens from its id
func (k Keeper) GetTokens(ctx context.Context, id uint64) (val types.Tokens, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokensKey))
	b := store.Get(GetTokensIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTokens removes a tokens from the store
func (k Keeper) RemoveTokens(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokensKey))
	store.Delete(GetTokensIDBytes(id))
}

// GetAllTokens returns all tokens
func (k Keeper) GetAllTokens(ctx context.Context) (list []types.Tokens) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TokensKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tokens
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTokensIDBytes returns the byte representation of the ID
func GetTokensIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.TokensKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
