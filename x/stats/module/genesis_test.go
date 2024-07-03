/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package stats_test

import (
	"testing"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	stats "github.com/stalwart-algoritmiclab/stwart-chain-go/x/stats/module"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/stats/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FeeStatsList: []types.FeeStats{
			{
				Date: "0",
			},
			{
				Date: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StatsKeeper(t)
	stats.InitGenesis(ctx, k, genesisState)
	got := stats.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FeeStatsList, got.FeeStatsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
