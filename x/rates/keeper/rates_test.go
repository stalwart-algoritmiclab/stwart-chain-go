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
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRates(keeper keeper.Keeper, ctx context.Context, n int) []types.Rates {
	items := make([]types.Rates, n)
	for i := range items {
		items[i].Denom = strconv.Itoa(i)

		keeper.SetRates(ctx, items[i])
	}
	return items
}

func TestRatesGet(t *testing.T) {
	keeper, ctx := keepertest.RatesKeeper(t)
	items := createNRates(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRates(ctx,
			item.Denom,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestRatesRemove(t *testing.T) {
	keeper, ctx := keepertest.RatesKeeper(t)
	items := createNRates(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveRates(ctx,
			item.Denom,
		)
		_, found := keeper.GetRates(ctx,
			item.Denom,
		)
		require.False(t, found)
	}
}

func TestRatesGetAll(t *testing.T) {
	keeper, ctx := keepertest.RatesKeeper(t)
	items := createNRates(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRates(ctx)),
	)
}
