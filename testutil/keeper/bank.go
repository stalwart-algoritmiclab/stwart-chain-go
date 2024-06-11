/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	addresscodec "github.com/cosmos/cosmos-sdk/codec/address"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/std"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/rand"
	coretypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
)

func BankKeeperWithAccounts(
	t testing.TB,
	countAccounts int,
	denoms ...string,
) (keeper.Keeper, sdk.Context, []simtypes.Account) {
	storeKey := storetypes.NewKVStoreKey(banktypes.StoreKey)
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(PrefixSTWART, "pub")

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())
	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authtypes.RegisterInterfaces(registry)
	std.RegisterInterfaces(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	storeService := runtime.NewKVStoreService(storeKey)

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		storeService,
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		authority.String(),
	)

	k := keeper.NewBaseKeeper(
		cdc,
		storeService,
		accountKeeper,
		BlockedAddresses(),
		authority.String(),
		log.NewNopLogger(),
	)

	// Initialize params
	if err := k.SetParams(ctx, banktypes.DefaultParams()); err != nil {
		panic(err)
	}

	// generate random accounts
	accounts := simtypes.RandomAccounts(rand.NewRand(), countAccounts)
	if len(accounts) < countAccounts {
		t.Errorf("must have at least %d accounts", countAccounts)
	}

	// create slices to hold coins for module and accounts
	moduleCoins := make(sdk.Coins, 0, len(denoms))
	accountCoins := make(sdk.Coins, 0, len(denoms))

	// generate coins for each denom
	for _, denom := range denoms {
		// the balance should not be empty after the sending operation, but equal to the amount of all users
		moduleAmount := int64(2 * countAccounts * testAmount)

		moduleCoins = append(moduleCoins, sdk.NewCoin(denom, math.NewInt(moduleAmount)))
		accountCoins = append(accountCoins, sdk.NewCoin(denom, math.NewInt(testAmount)))
	}

	// mint coins to module account
	if err := k.MintCoins(ctx, coretypes.ModuleName, moduleCoins); err != nil {
		panic(err)
	}

	// send coins from module to accounts.
	for _, a := range accounts {
		if err := k.SendCoinsFromModuleToAccount(ctx, coretypes.ModuleName, a.Address, accountCoins); err != nil {
			panic(err)
		}
	}

	return k, ctx, accounts
}
