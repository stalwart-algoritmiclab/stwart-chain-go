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

func createNPolls(keeper keeper.Keeper, ctx context.Context, n int) []types.Polls {
	items := make([]types.Polls, n)
	for i := range items {
		items[i].Id = keeper.AppendPolls(ctx, items[i])
	}
	return items
}

func TestPollsGet(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	items := createNPolls(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetPolls(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestPollsRemove(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	items := createNPolls(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemovePolls(ctx, item.Id)
		_, found := keeper.GetPolls(ctx, item.Id)
		require.False(t, found)
	}
}

func TestPollsGetAll(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	items := createNPolls(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllPolls(ctx)),
	)
}

func TestPollsCount(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	items := createNPolls(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetPollsCount(ctx))
}
