/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				PollsParams: &types.PollsParams{
					ProposerDeposit: 63,
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated votes",
			genState: &types.GenesisState{
				VotesList: []types.Votes{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid votes count",
			genState: &types.GenesisState{
				VotesList: []types.Votes{
					{
						Id: 1,
					},
				},
				VotesCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated options",
			genState: &types.GenesisState{
				OptionsList: []types.Options{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid options count",
			genState: &types.GenesisState{
				OptionsList: []types.Options{
					{
						Id: 1,
					},
				},
				OptionsCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated polls",
			genState: &types.GenesisState{
				PollsList: []types.Polls{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid polls count",
			genState: &types.GenesisState{
				PollsList: []types.Polls{
					{
						Id: 1,
					},
				},
				PollsCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
