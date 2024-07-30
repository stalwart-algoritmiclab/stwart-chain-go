/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package polls

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/sample"
	pollssimulation "github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/simulation"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/polls/types"
)

// avoid unused import issue
var (
	_ = pollssimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreatePollsParams = "op_weight_msg_polls_params"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePollsParams int = 100

	opWeightMsgUpdatePollsParams = "op_weight_msg_polls_params"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdatePollsParams int = 100

	opWeightMsgDeletePollsParams = "op_weight_msg_polls_params"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeletePollsParams int = 100

	opWeightMsgCreatePoll = "op_weight_msg_create_poll"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreatePoll int = 100

	opWeightMsgVote = "op_weight_msg_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVote int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	pollsGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&pollsGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreatePollsParams int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePollsParams, &weightMsgCreatePollsParams, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePollsParams = defaultWeightMsgCreatePollsParams
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePollsParams,
		pollssimulation.SimulateMsgCreatePollsParams(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdatePollsParams int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdatePollsParams, &weightMsgUpdatePollsParams, nil,
		func(_ *rand.Rand) {
			weightMsgUpdatePollsParams = defaultWeightMsgUpdatePollsParams
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdatePollsParams,
		pollssimulation.SimulateMsgUpdatePollsParams(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeletePollsParams int
	simState.AppParams.GetOrGenerate(opWeightMsgDeletePollsParams, &weightMsgDeletePollsParams, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePollsParams = defaultWeightMsgDeletePollsParams
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeletePollsParams,
		pollssimulation.SimulateMsgDeletePollsParams(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreatePoll int
	simState.AppParams.GetOrGenerate(opWeightMsgCreatePoll, &weightMsgCreatePoll, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePoll = defaultWeightMsgCreatePoll
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreatePoll,
		pollssimulation.SimulateMsgCreatePoll(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVote int
	simState.AppParams.GetOrGenerate(opWeightMsgVote, &weightMsgVote, nil,
		func(_ *rand.Rand) {
			weightMsgVote = defaultWeightMsgVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVote,
		pollssimulation.SimulateMsgVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePollsParams,
			defaultWeightMsgCreatePollsParams,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				pollssimulation.SimulateMsgCreatePollsParams(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdatePollsParams,
			defaultWeightMsgUpdatePollsParams,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				pollssimulation.SimulateMsgUpdatePollsParams(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeletePollsParams,
			defaultWeightMsgDeletePollsParams,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				pollssimulation.SimulateMsgDeletePollsParams(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreatePoll,
			defaultWeightMsgCreatePoll,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				pollssimulation.SimulateMsgCreatePoll(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgVote,
			defaultWeightMsgVote,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				pollssimulation.SimulateMsgVote(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
