/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/domain"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

func (k Keeper) GetFees(ctx sdk.Context, receiverAssets sdk.Coins, sendAmount sdk.Coin) (val types.Fees, minRefAmount sdk.Coin, found bool) {
	minRefAmount = sdk.NewCoin(domain.DenomStake, math.ZeroInt())
	// Retrieve tariffs for the given asset's denomination
	tariffs, found := k.GetTariffs(ctx, sendAmount.Denom)
	if !found {
		// If no tariffs defined, return false
		return val, minRefAmount, false
	}

	if len(tariffs.Tariffs) == 0 {
		// If no tariffs defined, return false
		return val, minRefAmount, false
	}

	var fees []*types.Fees
	lastFoundTriggerAmount := math.ZeroInt()
	// Traverse through all defined tariffs
	for _, tariff := range tariffs.Tariffs {
		// Convert amount to integer. If conversion fails, continue to the next tariff
		triggerAmount, ok := math.NewIntFromString(tariff.GetAmount())
		if !ok {
			continue
		}

		minRefAmountRaw, ok := math.NewIntFromString(tariff.GetMinRefBalance())
		if !ok {
			continue
		}

		// If tariff amount is zero, consider it as a default tariff and return the fees
		if triggerAmount.IsZero() {
			fees = tariff.Fees
			minRefAmount = sdk.NewCoin(tariff.Denom, minRefAmountRaw)
		}

		// Check if the coin denomination exists in sender's assets
		found, coin := receiverAssets.Find(tariff.Denom)
		if !found {
			// If default tariff is not defined, continue to the next tariff
			continue
		}

		// select fees for the matching tariff
		if coin.Amount.GTE(triggerAmount) {
			if lastFoundTriggerAmount.IsZero() || triggerAmount.GTE(lastFoundTriggerAmount) {
				fees = tariff.Fees
				lastFoundTriggerAmount = triggerAmount

				minRefAmount = sdk.NewCoin(tariff.Denom, minRefAmountRaw)
			}
		}
	}

	lastFoundTriggerAmount = math.ZeroInt()
	// Traverse through all defined tariffs
	for _, fee := range fees {
		// Convert amount to integer. If conversion fails, continue to the next tariff
		triggerAmount, ok := math.NewIntFromString(fee.GetAmountFrom())
		if !ok {
			continue
		}
		// If the sender has assets greater than or equal to the trigger amount, assign the fee value
		if sendAmount.Amount.GTE(triggerAmount) {
			if lastFoundTriggerAmount.IsZero() || triggerAmount.GTE(lastFoundTriggerAmount) {
				val = *fee
				lastFoundTriggerAmount = triggerAmount
			}
		}
	}

	// If no tariff was found, return false
	if val.Fee == "" {
		return val, minRefAmount, false
	}

	// Fees for the matching tariff is found
	return val, minRefAmount, true
}

// SetTariffs set a specific tariffs in the store from its index
func (k Keeper) SetTariffs(ctx context.Context, tariffs types.Tariffs) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TariffsKeyPrefix))
	b := k.cdc.MustMarshal(&tariffs)
	store.Set(types.TariffsKey(
		tariffs.Denom,
	), b)
}

// GetTariffs returns a tariffs from its index
func (k Keeper) GetTariffs(
	ctx context.Context,
	denom string,

) (val types.Tariffs, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TariffsKeyPrefix))

	b := store.Get(types.TariffsKey(
		denom,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTariffs removes a tariffs from the store
func (k Keeper) RemoveTariffs(
	ctx context.Context,
	denom string,

) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TariffsKeyPrefix))
	store.Delete(types.TariffsKey(
		denom,
	))
}

// GetAllTariffs returns all tariffs
func (k Keeper) GetAllTariffs(ctx context.Context) (list []types.Tariffs) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TariffsKeyPrefix))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Tariffs
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
