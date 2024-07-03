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

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNTariffs(keeper keeper.Keeper, ctx context.Context, n int) []types.Tariffs {
	items := make([]types.Tariffs, n)
	for i := range items {
		items[i].Denom = strconv.Itoa(i)

		keeper.SetTariffs(ctx, items[i])
	}
	return items
}

func TestTariffsGet(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNTariffs(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetTariffs(ctx,
			item.Denom,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestTariffsRemove(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNTariffs(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTariffs(ctx,
			item.Denom,
		)
		_, found := keeper.GetTariffs(ctx,
			item.Denom,
		)
		require.False(t, found)
	}
}

func TestTariffsGetAll(t *testing.T) {
	keeper, ctx := keepertest.FeepolicyKeeper(t)
	items := createNTariffs(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTariffs(ctx)),
	)
}
