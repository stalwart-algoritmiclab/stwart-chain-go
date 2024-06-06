/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper

import (
	"context"
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
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/exchanger/types"
	rateskeeper "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/keeper"
	ratestypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/rates/types"
)

const (
	testAddress = "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"
)

func ExchangerKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount("stwart", "pub")
	config.Seal()

	db := dbm.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, log.NewNopLogger(), metrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	require.NoError(t, stateStore.LoadLatestVersion())
	registry := codectypes.NewInterfaceRegistry()
	auth.RegisterInterfaces(registry)
	std.RegisterInterfaces(registry)
	cdc := codec.NewProtoCodec(registry)
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)

	securedKeeper, _ := SecuredKeeper(t) // unused keeper

	storeService := runtime.NewKVStoreService(storeKey)
	ratesKeeper := rateskeeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
		securedKeeper,
	)

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		storeService,
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec("stwart"),
		"stwart",
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		storeService,
		accountKeeper,
		nil, // TODO: previously app.BlockedModuleAccountAddrs()
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
		log.NewNopLogger(),
	)

	k := keeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(storeKey),
		log.NewNopLogger(),
		authority.String(),
		bankKeeper,
		accountKeeper,
		ratesKeeper,
	)

	ctx := sdk.NewContext(stateStore, cmtproto.Header{}, false, log.NewNopLogger())

	if err := prepareTestData(ctx, bankKeeper, ratesKeeper); err != nil {
		require.NoError(t, err)
	}

	// Initialize params
	if err := k.SetParams(ctx, types.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx
}

func prepareTestData(ctx context.Context, bankKeeper bankkeeper.Keeper, ratesKeeper rateskeeper.Keeper) error {
	if err := bankKeeper.MintCoins(ctx, types.ModuleName, sdk.Coins{
		sdk.Coin{
			Denom:  "ssc",
			Amount: math.NewInt(100_0000_0000),
		},
		sdk.Coin{
			Denom:  "usdt",
			Amount: math.NewInt(100_0000_0000),
		},
	}); err != nil {
		return err
	}

	if err := bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.MustAccAddressFromBech32(testAddress),
		sdk.Coins{
			sdk.Coin{
				Denom:  "ssc",
				Amount: math.NewInt(10_0000_0000),
			},
			sdk.Coin{
				Denom:  "usdt",
				Amount: math.NewInt(10_0000_0000),
			},
		},
	); err != nil {
		return err
	}

	ratesKeeper.SetRates(ctx, ratestypes.Rates{
		Denom:    "usdt",
		Rate:     1,
		Creator:  testAddress,
		Decimals: 8,
	})

	return nil
}
