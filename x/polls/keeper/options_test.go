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
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

func createNOptions(keeper keeper.Keeper, ctx context.Context, n int) []types.Options {
	items := make([]types.Options, n)
	for i := range items {
		items[i].Id = keeper.AppendOptions(ctx, items[i])
	}
	return items
}

func TestOptionsGet(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	items := createNOptions(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetOptions(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestOptionsRemove(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	items := createNOptions(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveOptions(ctx, item.Id)
		_, found := keeper.GetOptions(ctx, item.Id)
		require.False(t, found)
	}
}

func TestOptionsGetAll(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	items := createNOptions(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllOptions(ctx)),
	)
}

func TestOptionsCount(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	items := createNOptions(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetOptionsCount(ctx))
}
