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

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"
	securedkeeper "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/keeper"
	securedtypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"
)

const (
	feePolicyTestAccount  = "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at"
	feePolicyTestAccount2 = "stwart1tygms3xhhs3yv487phx3dw4a95jn7t7lqaz522"
)

func FeepolicyKeeper(t testing.TB) (keeper.Keeper, sdk.Context) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)

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

	prepareFeePolicyTestData(ctx, securedKeeper, k)

	// Initialize params
	if err := k.SetParams(ctx, types.DefaultParams()); err != nil {
		panic(err)
	}

	return k, ctx
}

func prepareFeePolicyTestData(ctx context.Context, securedkeeper securedkeeper.Keeper, feepolicyKeeper keeper.Keeper) {
	securedkeeper.AppendAddresses(ctx, securedtypes.Addresses{
		Id:      1,
		Address: []string{feePolicyTestAccount},
		Creator: feePolicyTestAccount,
	})

	feepolicyKeeper.AppendAddresses(ctx, types.Address{
		Id:      5,
		Address: feePolicyTestAccount2,
		Creator: feePolicyTestAccount2,
	})

	feepolicyKeeper.SetTariffs(ctx, types.Tariffs{
		Denom: "ccs",
		Tariffs: []*types.Tariff{
			{
				Denom:         "ccs",
				Id:            5,
				Amount:        "100",
				MinRefBalance: "100",
				Fees: []*types.Fees{
					{
						AmountFrom:  "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
						Fee:         "100",
						RefReward:   "100",
						StakeReward: "100",
						MinAmount:   100,
						NoRefReward: true,
						Creator:     "stwart1hdl6ny2kdpvth9p7u43ar9qer7tcvualelp0at",
						Id:          5,
					},
				},
			},
		},
	},
	)

	feepolicyKeeper.SetTariffs(ctx, types.Tariffs{
		Denom: "ccss",
		Tariffs: []*types.Tariff{
			{
				Denom:         "ccss",
				Id:            6,
				Amount:        "100",
				MinRefBalance: "100",
				Fees:          []*types.Fees{},
			},
		},
	},
	)
}
