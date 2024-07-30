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

func createTestPollsParams(keeper keeper.Keeper, ctx context.Context) types.PollsParams {
	item := types.PollsParams{}
	keeper.SetPollsParams(ctx, item)
	return item
}

func TestPollsParamsGet(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	item := createTestPollsParams(keeper, ctx)
	rst, found := keeper.GetPollsParams(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPollsParamsRemove(t *testing.T) {
	keeper, ctx := keepertest.PollsKeeper(t)
	createTestPollsParams(keeper, ctx)
	keeper.RemovePollsParams(ctx)
	_, found := keeper.GetPollsParams(ctx)
	require.False(t, found)
}
