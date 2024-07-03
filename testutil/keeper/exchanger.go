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
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/stretchr/testify/require"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/domain"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/exchanger/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/exchanger/types"
	rateskeeper "github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/keeper"
	ratestypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"
	securedkeeper "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/keeper"
)

func ExchangerKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(PrefixSTWART, "pub")
	config.Seal()

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

	securedKeeper := securedkeeper.NewKeeper(
		cdc,
		storeService,
		log.NewNopLogger(),
		authority.String(),
	)

	ratesKeeper := rateskeeper.NewKeeper(
		cdc,
		storeService,
		log.NewNopLogger(),
		authority.String(),
		securedKeeper,
	)

	accountKeeper := authkeeper.NewAccountKeeper(
		cdc,
		storeService,
		authtypes.ProtoBaseAccount,
		maccPerms,
		addresscodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	bankKeeper := bankkeeper.NewBaseKeeper(
		cdc,
		storeService,
		accountKeeper,
		BlockedAddresses(),
		authority.String(),
		log.NewNopLogger(),
	)

	k := keeper.NewKeeper(
		cdc,
		storeService,
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
			Denom:  domain.DenomStableIndex,
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
				Denom:  domain.DenomStableIndex,
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
