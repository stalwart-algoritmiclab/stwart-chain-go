/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper_test

import (
	"context"
	"testing"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/rand"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/faucet/types"
	securedtypes "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
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
