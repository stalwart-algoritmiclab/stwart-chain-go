/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package referral_test

import (
	"testing"

	keepertest "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/nullify"
	referral "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/module"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/referral/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		UserList: []types.User{
			{
				AccountAddress: "0",
			},
			{
				AccountAddress: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ReferralKeeper(t)
	referral.InitGenesis(ctx, k, genesisState)
	got := referral.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.UserList, got.UserList)
	// this line is used by starport scaffolding # genesis/test/assert
}
