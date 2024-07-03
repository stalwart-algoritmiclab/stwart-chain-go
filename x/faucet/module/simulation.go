/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package faucet

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/sample"
	faucetsimulation "github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/simulation"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/faucet/types"
)

// avoid unused import issue
var (
	_ = faucetsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateTokens = "op_weight_msg_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTokens int = 100

	opWeightMsgUpdateTokens = "op_weight_msg_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTokens int = 100

	opWeightMsgDeleteTokens = "op_weight_msg_tokens"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTokens int = 100

	opWeightMsgIssue = "op_weight_msg_issue"
	// TODO: Determine the simulation weight value
	defaultWeightMsgIssue int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	faucetGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TokensList: []types.Tokens{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TokensCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&faucetGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTokens int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTokens, &weightMsgCreateTokens, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTokens = defaultWeightMsgCreateTokens
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTokens,
		faucetsimulation.SimulateMsgCreateTokens(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTokens int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateTokens, &weightMsgUpdateTokens, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTokens = defaultWeightMsgUpdateTokens
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTokens,
		faucetsimulation.SimulateMsgUpdateTokens(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTokens int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteTokens, &weightMsgDeleteTokens, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTokens = defaultWeightMsgDeleteTokens
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTokens,
		faucetsimulation.SimulateMsgDeleteTokens(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgIssue int
	simState.AppParams.GetOrGenerate(opWeightMsgIssue, &weightMsgIssue, nil,
		func(_ *rand.Rand) {
			weightMsgIssue = defaultWeightMsgIssue
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgIssue,
		faucetsimulation.SimulateMsgIssue(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTokens,
			defaultWeightMsgCreateTokens,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				faucetsimulation.SimulateMsgCreateTokens(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTokens,
			defaultWeightMsgUpdateTokens,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				faucetsimulation.SimulateMsgUpdateTokens(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteTokens,
			defaultWeightMsgDeleteTokens,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				faucetsimulation.SimulateMsgDeleteTokens(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgIssue,
			defaultWeightMsgIssue,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				faucetsimulation.SimulateMsgIssue(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
