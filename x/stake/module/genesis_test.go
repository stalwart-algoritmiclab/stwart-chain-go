package stake_test

import (
	"testing"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	stake "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stake/module"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/stake/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		StakeList: []types.Stake{
			{
				Address: "0",
			},
			{
				Address: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.StakeKeeper(t)
	stake.InitGenesis(ctx, k, genesisState)
	got := stake.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.StakeList, got.StakeList)
	// this line is used by starport scaffolding # genesis/test/assert
}
