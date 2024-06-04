package types_test

import (
	"testing"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"

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