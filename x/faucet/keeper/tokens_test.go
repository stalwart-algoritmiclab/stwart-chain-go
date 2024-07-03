/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/types"
)

func createNTokens(keeper keeper.Keeper, ctx context.Context, n int) []types.Tokens {
	items := make([]types.Tokens, n)
	for i := range items {
		items[i].Id = keeper.AppendTokens(ctx, items[i])
	}
	return items
}

func TestTokensGet(t *testing.T) {
	keeper, ctx := keepertest.FaucetKeeper(t)
	items := createNTokens(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTokens(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTokensRemove(t *testing.T) {
	keeper, ctx := keepertest.FaucetKeeper(t)
	items := createNTokens(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTokens(ctx, item.Id)
		_, found := keeper.GetTokens(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTokensGetAll(t *testing.T) {
	keeper, ctx := keepertest.FaucetKeeper(t)
	items := createNTokens(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTokens(ctx)),
	)
}

func TestTokensCount(t *testing.T) {
	keeper, ctx := keepertest.FaucetKeeper(t)
	items := createNTokens(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTokensCount(ctx))
}
