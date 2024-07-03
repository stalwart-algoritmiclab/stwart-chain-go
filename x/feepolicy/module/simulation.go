/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package feepolicy

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/stalwart-algoritmiclab/stwart-chain-go/testutil/sample"
	feepolicysimulation "github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/simulation"
	"github.com/stalwart-algoritmiclab/stwart-chain-go/x/feepolicy/types"
)

// avoid unused import issue
var (
	_ = feepolicysimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgCreateAddresses = "op_weight_msg_addresses"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAddresses int = 100

	opWeightMsgUpdateAddresses = "op_weight_msg_addresses"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAddresses int = 100

	opWeightMsgDeleteAddresses = "op_weight_msg_addresses"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAddresses int = 100

	opWeightMsgCreateTariffs = "op_weight_msg_tariffs"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTariffs int = 100

	opWeightMsgUpdateTariffs = "op_weight_msg_tariffs"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTariffs int = 100

	opWeightMsgDeleteTariffs = "op_weight_msg_tariffs"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTariffs int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	feepolicyGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AddressesList: []types.Address{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		AddressesCount: 2,
		TariffsList: []types.Tariffs{
			{
				Creator: sample.AccAddress(),
				Denom:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Denom:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&feepolicyGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAddresses int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateAddresses, &weightMsgCreateAddresses, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAddresses = defaultWeightMsgCreateAddresses
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAddresses,
		feepolicysimulation.SimulateMsgCreateAddresses(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAddresses int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateAddresses, &weightMsgUpdateAddresses, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAddresses = defaultWeightMsgUpdateAddresses
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAddresses,
		feepolicysimulation.SimulateMsgUpdateAddresses(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAddresses int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteAddresses, &weightMsgDeleteAddresses, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAddresses = defaultWeightMsgDeleteAddresses
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAddresses,
		feepolicysimulation.SimulateMsgDeleteAddresses(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTariffs int
	simState.AppParams.GetOrGenerate(opWeightMsgCreateTariffs, &weightMsgCreateTariffs, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTariffs = defaultWeightMsgCreateTariffs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTariffs,
		feepolicysimulation.SimulateMsgCreateTariffs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTariffs int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateTariffs, &weightMsgUpdateTariffs, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTariffs = defaultWeightMsgUpdateTariffs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTariffs,
		feepolicysimulation.SimulateMsgUpdateTariffs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTariffs int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteTariffs, &weightMsgDeleteTariffs, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTariffs = defaultWeightMsgDeleteTariffs
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTariffs,
		feepolicysimulation.SimulateMsgDeleteTariffs(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAddresses,
			defaultWeightMsgCreateAddresses,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				feepolicysimulation.SimulateMsgCreateAddresses(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAddresses,
			defaultWeightMsgUpdateAddresses,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				feepolicysimulation.SimulateMsgUpdateAddresses(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAddresses,
			defaultWeightMsgDeleteAddresses,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				feepolicysimulation.SimulateMsgDeleteAddresses(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTariffs,
			defaultWeightMsgCreateTariffs,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				feepolicysimulation.SimulateMsgCreateTariffs(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTariffs,
			defaultWeightMsgUpdateTariffs,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				feepolicysimulation.SimulateMsgUpdateTariffs(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteTariffs,
			defaultWeightMsgDeleteTariffs,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				feepolicysimulation.SimulateMsgDeleteTariffs(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
