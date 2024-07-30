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

type PollStatus string

// Block which defines all possible status types for Issue.
const (
	StatusPending      PollStatus = "pending"
	StatusVotingPeriod PollStatus = "voting_period"
	StatusPassed       PollStatus = "passed"
	StatusRejected     PollStatus = "rejected"
	StatusFailed       PollStatus = "failed"
)

// String - method for casting status to string.
func (s PollStatus) String() string {
	return string(s)
}

// GetPollsCount get the total number of polls
func (k Keeper) GetPollsCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PollsCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetPollsCount set the total number of polls
func (k Keeper) SetPollsCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.PollsCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendPolls appends a polls in the store with a new id and update the count
func (k Keeper) AppendPolls(
	ctx context.Context,
	polls types.Polls,
) uint64 {
	// Create the polls
	count := k.GetPollsCount(ctx)

	// Set the ID of the appended value
	polls.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PollsKey))
	appendedValue := k.cdc.MustMarshal(&polls)
	store.Set(GetPollsIDBytes(polls.Id), appendedValue)

	// Update polls count
	k.SetPollsCount(ctx, count+1)

	return count
}

// SetPolls set a specific polls in the store
func (k Keeper) SetPolls(ctx context.Context, polls types.Polls) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PollsKey))
	b := k.cdc.MustMarshal(&polls)
	store.Set(GetPollsIDBytes(polls.Id), b)
}

// GetPolls returns a polls from its id
func (k Keeper) GetPolls(ctx context.Context, id uint64) (val types.Polls, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PollsKey))
	b := store.Get(GetPollsIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetPollsOptionByID returns a polls option from its id
func (k Keeper) GetPollsOptionByID(poll types.Polls, id uint64) (*types.Options, bool) {
	for _, option := range poll.Options {
		if option.Id == id {
			return option, true
		}
	}

	return nil, false
}

// UpdateOption updates a polls option from its id
func UpdateOption(poll types.Polls, option *types.Options) (types.Polls, bool) {
	for idx, opt := range poll.Options {
		if opt.Id == option.Id {
			poll.Options[idx] = option
			return poll, true
		}
	}

	return poll, false
}

// RemovePolls removes a polls from the store
func (k Keeper) RemovePolls(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PollsKey))
	store.Delete(GetPollsIDBytes(id))
}

// GetAllPolls returns all polls
func (k Keeper) GetAllPolls(ctx context.Context) (list []types.Polls) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PollsKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Polls
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAllPollsByStatuses returns all polls by statuses
func (k Keeper) GetAllPollsByStatuses(ctx context.Context, statuses ...PollStatus) (list []types.Polls) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.PollsKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	statusMap := map[string]bool{}
	for _, status := range statuses {
		statusMap[status.String()] = true
	}

	for ; iterator.Valid(); iterator.Next() {
		var val types.Polls
		k.cdc.MustUnmarshal(iterator.Value(), &val)

		if _, exists := statusMap[val.Status]; exists {
			list = append(list, val)
		}
	}

	return
}

// GetPollsIDBytes returns the byte representation of the ID
func GetPollsIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.PollsKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
