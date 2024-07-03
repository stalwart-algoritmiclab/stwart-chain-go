/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/stake/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetStake set a specific stake in the store from its index
func (k Keeper) SetStake(ctx context.Context, stake types.Stake) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StakeKeyPrefix))
	b := k.cdc.MustMarshal(&stake)
	store.Set(types.StakeKey(
		stake.Address,
	), b)
}

// GetStake returns a stake from its index
func (k Keeper) GetStake(ctx context.Context, address string) (val types.Stake, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StakeKeyPrefix))

	b := store.Get(types.StakeKey(
		address,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// GetFreeStake returns a stake amount without stake sell
func (k Keeper) GetFreeStake(ctx context.Context, address sdk.AccAddress) sdkmath.Int {
	assets := k.bankKeeper.SpendableCoins(ctx, address)
	spendable := assets.AmountOf(domain.DenomStake)
	if spendable.IsPositive() {
		userStake, found := k.GetStake(ctx, address.String())
		if !found {
			return spendable
		}

		freeStakeAmount := spendable.Sub(userStake.SellAmount.Amount)
		if freeStakeAmount.IsPositive() {
			return freeStakeAmount
		}
	}

	return sdkmath.NewInt(0)
}

// RemoveStake removes a stake from the store
func (k Keeper) RemoveStake(ctx context.Context, address string) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StakeKeyPrefix))
	store.Delete(types.StakeKey(
		address,
	))
}

// GetAllStake returns all stake
func (k Keeper) GetAllStake(ctx context.Context) (list []types.Stake) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.StakeKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Stake
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
