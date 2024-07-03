/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package faucet_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	faucet "github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/module"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		TokensList: []types.Tokens{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TokensCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FaucetKeeper(t)
	faucet.InitGenesis(ctx, k, genesisState)
	got := faucet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.TokensList, got.TokensList)
	require.Equal(t, genesisState.TokensCount, got.TokensCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
