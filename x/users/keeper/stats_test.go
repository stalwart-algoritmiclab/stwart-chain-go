/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/users/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStats(keeper keeper.Keeper, ctx context.Context, n int) []types.Stats {
	items := make([]types.Stats, n)
	for i := range items {
		items[i].Date = strconv.Itoa(i)

		keeper.SetStats(ctx, items[i])
	}
	return items
}

func TestStatsGet(t *testing.T) {
	keeper, ctx := keepertest.UsersKeeper(t)
	items := createNStats(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStats(ctx,
			item.Date,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStatsRemove(t *testing.T) {
	keeper, ctx := keepertest.UsersKeeper(t)
	items := createNStats(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStats(ctx,
			item.Date,
		)
		_, found := keeper.GetStats(ctx,
			item.Date,
		)
		require.False(t, found)
	}
}

func TestStatsGetAll(t *testing.T) {
	keeper, ctx := keepertest.UsersKeeper(t)
	items := createNStats(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStats(ctx)),
	)
}
