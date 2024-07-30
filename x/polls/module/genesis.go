/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package polls

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/keeper"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.PollsParams != nil {
		k.SetPollsParams(ctx, *genState.PollsParams)
	}
	// Set all the votes
	for _, elem := range genState.VotesList {
		k.SetVotes(ctx, elem)
	}

	// Set votes count
	k.SetVotesCount(ctx, genState.VotesCount)
	// Set all the options
	for _, elem := range genState.OptionsList {
		k.SetOptions(ctx, elem)
	}

	// Set options count
	k.SetOptionsCount(ctx, genState.OptionsCount)
	// Set all the polls
	for _, elem := range genState.PollsList {
		k.SetPolls(ctx, elem)
	}

	// Set polls count
	k.SetPollsCount(ctx, genState.PollsCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all pollsParams
	pollsParams, found := k.GetPollsParams(ctx)
	if found {
		genesis.PollsParams = &pollsParams
	}
	genesis.VotesList = k.GetAllVotes(ctx)
	genesis.VotesCount = k.GetVotesCount(ctx)
	genesis.OptionsList = k.GetAllOptions(ctx)
	genesis.OptionsCount = k.GetOptionsCount(ctx)
	genesis.PollsList = k.GetAllPolls(ctx)
	genesis.PollsCount = k.GetPollsCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
