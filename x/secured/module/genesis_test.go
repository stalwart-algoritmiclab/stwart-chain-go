/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package secured_test

import (
	"testing"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	secured "github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/module"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/secured/types"

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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SecuredKeeper(t)
	secured.InitGenesis(ctx, k, genesisState)
	got := secured.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AddressesList, got.AddressesList)
	require.Equal(t, genesisState.AddressesCount, got.AddressesCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
