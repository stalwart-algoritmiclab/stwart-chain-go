/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package referral

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/sample"
	referralsimulation "github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/simulation"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/referral/types"
)

// avoid unused import issue
var (
	_ = referralsimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateUser = "op_weight_msg_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateUser int = 100

	opWeightMsgUpdateUser = "op_weight_msg_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateUser int = 100

	opWeightMsgDeleteUser = "op_weight_msg_user"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteUser int = 100

	opWeightMsgSetReferrer = "op_weight_msg_set_referrer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSetReferrer int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	referralGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		UserList: []types.User{
			{
				AccountAddress: "0",
			},
			{
				AccountAddress: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&referralGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateUser int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateUser, &weightMsgCreateUser, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUser = defaultWeightMsgCreateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateUser,
		referralsimulation.SimulateMsgCreateUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateUser int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateUser, &weightMsgUpdateUser, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateUser = defaultWeightMsgUpdateUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateUser,
		referralsimulation.SimulateMsgUpdateUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteUser int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteUser, &weightMsgDeleteUser, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUser = defaultWeightMsgDeleteUser
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteUser,
		referralsimulation.SimulateMsgDeleteUser(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSetReferrer int
	simState.AppParams.GetOrGenerate(opWeightMsgSetReferrer, &weightMsgSetReferrer, nil,
		func(_ *rand.Rand) {
			weightMsgSetReferrer = defaultWeightMsgSetReferrer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSetReferrer,
		referralsimulation.SimulateMsgSetReferrer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateUser,
			defaultWeightMsgCreateUser,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				referralsimulation.SimulateMsgCreateUser(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateUser,
			defaultWeightMsgUpdateUser,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				referralsimulation.SimulateMsgUpdateUser(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteUser,
			defaultWeightMsgDeleteUser,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				referralsimulation.SimulateMsgDeleteUser(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSetReferrer,
			defaultWeightMsgSetReferrer,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				referralsimulation.SimulateMsgSetReferrer(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
