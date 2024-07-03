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
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/stake/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/stake/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStake(keeper keeper.Keeper, ctx context.Context, n int) []types.Stake {
	items := make([]types.Stake, n)
	for i := range items {
		items[i].Address = strconv.Itoa(i)

		keeper.SetStake(ctx, items[i])
	}
	return items
}

func TestStakeGet(t *testing.T) {
	keeper, ctx := keepertest.StakeKeeper(t)
	items := createNStake(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStake(ctx,
			item.Address,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStakeRemove(t *testing.T) {
	keeper, ctx := keepertest.StakeKeeper(t)
	items := createNStake(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStake(ctx,
			item.Address,
		)
		_, found := keeper.GetStake(ctx,
			item.Address,
		)
		require.False(t, found)
	}
}

func TestStakeGetAll(t *testing.T) {
	keeper, ctx := keepertest.StakeKeeper(t)
	items := createNStake(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStake(ctx)),
	)
}
