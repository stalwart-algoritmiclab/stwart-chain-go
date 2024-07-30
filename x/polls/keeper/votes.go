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

// GetVotesCount get the total number of votes
func (k Keeper) GetVotesCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.VotesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVotesCount set the total number of votes
func (k Keeper) SetVotesCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.VotesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVotes appends a votes in the store with a new id and update the count
func (k Keeper) AppendVotes(
	ctx context.Context,
	votes types.Votes,
) uint64 {
	// Create the votes
	count := k.GetVotesCount(ctx)

	// Set the ID of the appended value
	votes.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VotesKey))
	appendedValue := k.cdc.MustMarshal(&votes)
	store.Set(GetVotesIDBytes(votes.Id), appendedValue)

	// Update votes count
	k.SetVotesCount(ctx, count+1)

	return count
}

// SetVotes set a specific votes in the store
func (k Keeper) SetVotes(ctx context.Context, votes types.Votes) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VotesKey))
	b := k.cdc.MustMarshal(&votes)
	store.Set(GetVotesIDBytes(votes.Id), b)
}

// GetVotes returns a votes from its id
func (k Keeper) GetVotes(ctx context.Context, id uint64) (val types.Votes, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VotesKey))
	b := store.Get(GetVotesIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVotes removes a votes from the store
func (k Keeper) RemoveVotes(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VotesKey))
	store.Delete(GetVotesIDBytes(id))
}

// GetAllVotes returns all votes
func (k Keeper) GetAllVotes(ctx context.Context) (list []types.Votes) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VotesKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Votes
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVoteByAccountAddressAndPollID returns all votes by account address and poll id
func (k Keeper) GetVoteByAccountAddressAndPollID(ctx context.Context, accountAddress string, pollID uint64) (types.Votes, bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.VotesKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Votes
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		if val.AccountAddress == accountAddress && val.PollId == pollID {
			return val, true
		}
	}

	return types.Votes{}, false
}

// GetVotesIDBytes returns the byte representation of the ID
func GetVotesIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.VotesKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
