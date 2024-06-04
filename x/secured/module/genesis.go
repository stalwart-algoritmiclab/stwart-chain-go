package secured

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/secured/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the addresses
	for _, elem := range genState.AddressesList {
		k.SetAddresses(ctx, elem)
	}

	// Set addresses count
	k.SetAddressesCount(ctx, genState.AddressesCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AddressesList = k.GetAllAddresses(ctx)
	genesis.AddressesCount = k.GetAddressesCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
