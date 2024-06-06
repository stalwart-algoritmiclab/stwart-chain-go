/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package types_test

import (
	"testing"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"

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

				AddressesList: []types.Addresses{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated addresses",
			genState: &types.GenesisState{
				AddressesList: []types.Addresses{
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
			desc: "invalid addresses count",
			genState: &types.GenesisState{
				AddressesList: []types.Addresses{
					{
						Id: 1,
					},
				},
				AddressesCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated tariff",
			genState: &types.GenesisState{
				TariffList: []types.Tariff{
					{
						Denom: "0",
					},
					{
						Denom: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated tariffs",
			genState: &types.GenesisState{
				TariffsList: []types.Tariffs{
					{
						Denom: "0",
					},
					{
						Denom: "0",
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
