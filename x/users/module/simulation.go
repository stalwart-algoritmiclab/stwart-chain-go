/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package users

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/sample"
	userssimulation "github.com/stalwart-algoritmiclab/stwart-chain-go/x/users/simulation"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/users/types"
)

// avoid unused import issue
var (
	_ = userssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateStats = "op_weight_msg_stats"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateStats int = 100

	opWeightMsgUpdateStats = "op_weight_msg_stats"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateStats int = 100

	opWeightMsgDeleteStats = "op_weight_msg_stats"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteStats int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	usersGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		StatsList: []types.Stats{
			{
				Date: "0",
			},
			{
				Date: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&usersGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateStats int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateStats, &weightMsgCreateStats, nil,
		func(_ *rand.Rand) {
			weightMsgCreateStats = defaultWeightMsgCreateStats
		},
	)
	var weightMsgUpdateStats int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateStats, &weightMsgUpdateStats, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateStats = defaultWeightMsgUpdateStats
		},
	)

	var weightMsgDeleteStats int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteStats, &weightMsgDeleteStats, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteStats = defaultWeightMsgDeleteStats
		},
	)

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
