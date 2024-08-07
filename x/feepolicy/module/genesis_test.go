/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package feepolicy_test

import (
	"testing"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	feepolicy "github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/module"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AddressesList: []types.Address{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AddressesCount: 2,
		TariffList: []types.Tariff{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		TariffsList: []types.Tariffs{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FeepolicyKeeper(t)
	feepolicy.InitGenesis(ctx, k, genesisState)
	got := feepolicy.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AddressesList, got.AddressesList)
	require.Equal(t, genesisState.AddressesCount, got.AddressesCount)
	require.ElementsMatch(t, genesisState.TariffList, got.TariffList)
	require.ElementsMatch(t, genesisState.TariffsList, got.TariffsList)
	// this line is used by starport scaffolding # genesis/test/assert
}
