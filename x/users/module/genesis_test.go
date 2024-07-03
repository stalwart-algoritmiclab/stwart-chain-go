/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package users_test

import (
	"testing"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	users "github.com/stalwart-algoritmiclab/stwart-chain-go/x/users/module"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/users/types"

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
		UniqueUsersList: []types.UniqueUsers{
			{
				Date: "0",
			},
			{
				Date: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UsersKeeper(t)
	users.InitGenesis(ctx, k, genesisState)
	got := users.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StatsList, got.StatsList)
	require.ElementsMatch(t, genesisState.UniqueUsersList, got.UniqueUsersList)
	// this line is used by starport scaffolding # genesis/test/assert
}
