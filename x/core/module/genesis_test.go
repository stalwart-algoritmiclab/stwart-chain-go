/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package core_test

import (
	"testing"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	core "github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/module"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		StatsList: []types.Stats{
			{
				Date: "0",
			},
			{
				Date: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoreKeeper(t)
	core.InitGenesis(ctx, k, genesisState)
	got := core.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StatsList, got.StatsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
