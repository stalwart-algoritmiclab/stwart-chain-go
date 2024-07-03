/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package rates_test

import (
	"testing"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	rates "github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/module"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/rates/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AddressesList: []types.Addresses{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AddressesCount: 2,
		RatesList: []types.Rates{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RatesKeeper(t)
	rates.InitGenesis(ctx, k, genesisState)
	got := rates.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AddressesList, got.AddressesList)
	require.Equal(t, genesisState.AddressesCount, got.AddressesCount)
	require.ElementsMatch(t, genesisState.RatesList, got.RatesList)
	// this line is used by starport scaffolding # genesis/test/assert
}
