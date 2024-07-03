/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types_test

import (
	"testing"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/users/types"

	"github.com/stretchr/testify/require"
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated stats",
			genState: &types.GenesisState{
				StatsList: []types.Stats{
					{
						Date: "0",
					},
					{
						Date: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated uniqueUsers",
			genState: &types.GenesisState{
				UniqueUsersList: []types.UniqueUsers{
					{
						Date: "0",
					},
					{
						Date: "0",
					},
				},
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
