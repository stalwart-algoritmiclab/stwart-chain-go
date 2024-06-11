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
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/rand"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
	securedkeeper "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/keeper"
	securedtypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
	usersmodulekeeper "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/keeper"
)

func CoreKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(PrefixSTWART, "pub")
	// config.Seal()

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	authtypes.RegisterInterfaces(registry)
	std.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)
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

	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		storeService,
		accountKeeper,
		blockedModuleAccounts,
		authority.String(),
		log.NewNopLogger(),
	)

	securedKeeper := securedkeeper.NewKeeper(
		cdc,
		storeService,
		log.NewNopLogger(),
		authority.String(),
	)

	usersKeeper := usersmodulekeeper.NewKeeper(
		cdc,
		storeService,
		log.NewNopLogger(),
		authority.String(),
	)

	k := keeper.NewKeeper(
		cdc,
		storeService,
		log.NewNopLogger(),
		authority.String(),
		securedKeeper,
		usersKeeper,
		accountKeeper,
		bankKeeper,
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	// Initialize params
	if err := k.SetParams(ctx, types.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx
}

func CoreKeeperWithAddresses(
	t testing.TB,
	countAccounts int,
	denoms ...string,
) (keeper.Keeper, sdk.Context, []simtypes.Account) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(PrefixSTWART, "pub")
	// config.Seal()

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

	securedKeeper := securedkeeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
	)

	// generate random accounts
	accounts := simtypes.RandomAccounts(rand.NewRand(), countAccounts)
	addresses := make([]string, 0, len(accounts))
	for _, account := range accounts {
		addresses = append(addresses, account.Address.String())
	}

	creator := ""
	if len(addresses) != 0 {
		creator = accounts[0].Address.String()
	}

	// add addresses to keeper
	securedKeeper.AppendAddresses(ctx, securedtypes.Addresses{
		Id:      1,
		Address: addresses,
		Creator: creator,
	})

	usersKeeper := usersmodulekeeper.NewKeeper(
		cdc,
		storeService,
		log.NewNopLogger(),
		authority.String(),
	)

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		storeService,
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		authority.String(),
	)

	// add addresses to accountKeeper
	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		storeService,
		accountKeeper,
		BlockedAddresses(),
		authority.String(),
		log.NewNopLogger(),
	)

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
	if err := bankKeeper.MintCoins(ctx, types.ModuleName, moduleCoins); err != nil {
		panic(err)
	}

	// send coins from module to accounts.
	for _, a := range accounts {
		if err := bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, a.Address, accountCoins); err != nil {
			panic(err)
		}
	}

	k := keeper.NewKeeper(
		cdc,
		storeService,
		log.NewNopLogger(),
		authority.String(),
		securedKeeper,
		usersKeeper,
		accountKeeper,
		bankKeeper,
	)

	// Initialize params
	if err := k.SetParams(ctx, types.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx, accounts
}
