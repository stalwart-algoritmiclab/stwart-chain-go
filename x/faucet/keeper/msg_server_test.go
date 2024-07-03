/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"context"
	"testing"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/rand"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/types"
	securedtypes "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.FaucetKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func setupMsgServerWithAddresses(t testing.TB, countAccounts int) (
	keeper.Keeper,
	types.MsgServer,
	context.Context,
	[]simtypes.Account,
) {
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

	k, ctx := keepertest.FaucetKeeperWithAddresses(t, securedtypes.Addresses{
		Id:      1,
		Address: addresses,
		Creator: creator,
	})

	return k, keeper.NewMsgServerImpl(k), ctx, accounts
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
