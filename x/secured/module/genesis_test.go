package secured_test

import (
	"testing"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	secured "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/module"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"

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
