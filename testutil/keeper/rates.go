/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper

import (
	"context"
	"testing"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/keeper"
	ratestypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"
	securedkeeper "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/keeper"
	securedtypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"
)

const (
	testSecuredAddress = "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"
)

func RatesKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(ratestypes.StoreKey)

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	securedKeeper := securedkeeper.NewKeeper(cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String())

	k := keeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
		securedKeeper,
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	prepareRatesTestData(ctx, securedKeeper, k)

	// Initialize params
	if err := k.SetParams(ctx, ratestypes.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx
}

func prepareRatesTestData(ctx context.Context, securedKeeper securedkeeper.Keeper, ratesKeeper keeper.Keeper) {
	securedKeeper.AppendAddresses(ctx, securedtypes.Addresses{
		Address: []string{testSecuredAddress},
		Creator: testSecuredAddress,
	})

	ratesKeeper.SetRates(ctx, ratestypes.Rates{
		Denom:    "usdc",
		Rate:     1,
		Creator:  testSecuredAddress,
		Decimals: 8,
	})
}
