/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package polls_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/nullify"
	polls "github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/module"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PollsParams: &types.PollsParams{
			ProposerDeposit: 11,
			BurnVeto:        false,
		},
		VotesList: []types.Votes{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VotesCount: 2,
		OptionsList: []types.Options{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		OptionsCount: 2,
		PollsList: []types.Polls{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PollsCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PollsKeeper(t)
	polls.InitGenesis(ctx, k, genesisState)
	got := polls.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PollsParams, got.PollsParams)
	require.ElementsMatch(t, genesisState.VotesList, got.VotesList)
	require.Equal(t, genesisState.VotesCount, got.VotesCount)
	require.ElementsMatch(t, genesisState.OptionsList, got.OptionsList)
	require.Equal(t, genesisState.OptionsCount, got.OptionsCount)
	require.ElementsMatch(t, genesisState.PollsList, got.PollsList)
	require.Equal(t, genesisState.PollsCount, got.PollsCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
