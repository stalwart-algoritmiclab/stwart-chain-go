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

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/core/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.CoreKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}

func setupMsgServerWithAddresses(
	t testing.TB,
	countAccounts int,
	denoms ...string,
) (keeper.Keeper, types.MsgServer, context.Context, []simtypes.Account) {
	k, ctx, accounts := keepertest.CoreKeeperWithAddresses(t, countAccounts, denoms...)
	return k, keeper.NewMsgServerImpl(k), ctx, accounts
}
