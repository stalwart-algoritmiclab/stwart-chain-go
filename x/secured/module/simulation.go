package secured

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/testutil/sample"
	securedsimulation "gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/simulation"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
)

// avoid unused import issue
var (
	_ = securedsimulation.FindAccount
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

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	securedGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AddressesList: []types.Addresses{
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
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&securedGenesis)
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
		securedsimulation.SimulateMsgCreateAddresses(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAddresses int
	simState.AppParams.GetOrGenerate(opWeightMsgUpdateAddresses, &weightMsgUpdateAddresses, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAddresses = defaultWeightMsgUpdateAddresses
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAddresses,
		securedsimulation.SimulateMsgUpdateAddresses(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAddresses int
	simState.AppParams.GetOrGenerate(opWeightMsgDeleteAddresses, &weightMsgDeleteAddresses, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAddresses = defaultWeightMsgDeleteAddresses
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAddresses,
		securedsimulation.SimulateMsgDeleteAddresses(am.accountKeeper, am.bankKeeper, am.keeper),
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
				securedsimulation.SimulateMsgCreateAddresses(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAddresses,
			defaultWeightMsgUpdateAddresses,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				securedsimulation.SimulateMsgUpdateAddresses(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAddresses,
			defaultWeightMsgDeleteAddresses,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				securedsimulation.SimulateMsgDeleteAddresses(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
