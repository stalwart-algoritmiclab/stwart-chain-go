/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package core

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/sample"
	coresimulation "github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/simulation"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/core/types"
)

// avoid unused import issue
var (
	_ = coresimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgIssue = "op_weight_msg_issue"
	// TODO: Determine the simulation weight value
	defaultWeightMsgIssue int = 100

	opWeightMsgWithdraw = "op_weight_msg_withdraw"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdraw int = 100

	opWeightMsgSend = "op_weight_msg_send"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSend int = 100

	opWeightMsgRefund = "op_weight_msg_refund"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRefund int = 100

	opWeightMsgFees = "op_weight_msg_fees"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFees int = 100

	opWeightMsgRefReward = "op_weight_msg_ref_reward"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRefReward int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	coreGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&coreGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgIssue int
	simState.AppParams.GetOrGenerate(opWeightMsgIssue, &weightMsgIssue, nil,
		func(_ *rand.Rand) {
			weightMsgIssue = defaultWeightMsgIssue
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgIssue,
		coresimulation.SimulateMsgIssue(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdraw int
	simState.AppParams.GetOrGenerate(opWeightMsgWithdraw, &weightMsgWithdraw, nil,
		func(_ *rand.Rand) {
			weightMsgWithdraw = defaultWeightMsgWithdraw
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdraw,
		coresimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSend int
	simState.AppParams.GetOrGenerate(opWeightMsgSend, &weightMsgSend, nil,
		func(_ *rand.Rand) {
			weightMsgSend = defaultWeightMsgSend
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSend,
		coresimulation.SimulateMsgSend(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgIssue,
			defaultWeightMsgIssue,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				coresimulation.SimulateMsgIssue(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWithdraw,
			defaultWeightMsgWithdraw,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				coresimulation.SimulateMsgWithdraw(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSend,
			defaultWeightMsgSend,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				coresimulation.SimulateMsgSend(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
	}
}
