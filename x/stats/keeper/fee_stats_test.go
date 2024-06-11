/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stats/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stats/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFeeStats(keeper keeper.Keeper, ctx context.Context, n int) []types.FeeStats {
	items := make([]types.FeeStats, n)
	for i := range items {
		items[i].Date = strconv.Itoa(i)

		keeper.SetFeeStats(ctx, items[i])
	}
	return items
}

func TestFeeStatsGet(t *testing.T) {
	keeper, ctx := keepertest.StatsKeeper(t)
	items := createNFeeStats(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFeeStats(ctx,
			item.Date,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFeeStatsRemove(t *testing.T) {
	keeper, ctx := keepertest.StatsKeeper(t)
	items := createNFeeStats(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFeeStats(ctx,
			item.Date,
		)
		_, found := keeper.GetFeeStats(ctx,
			item.Date,
		)
		require.False(t, found)
	}
}

func TestFeeStatsGetAll(t *testing.T) {
	keeper, ctx := keepertest.StatsKeeper(t)
	items := createNFeeStats(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFeeStats(ctx)),
	)
}
