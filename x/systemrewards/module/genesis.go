/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package systemrewards

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/systemrewards/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/systemrewards/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the stats
	for _, elem := range genState.StatsList {
		k.SetStats(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.StatsList = k.GetAllStats(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
