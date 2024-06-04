package feepolicy

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/keeper"
	"gitlab.stalwart.tech/ijio/main/backend/stwart-chain/x/feepolicy/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the addresses
	for _, elem := range genState.AddressesList {
		k.SetAddresses(ctx, elem)
	}

	// Set addresses count
	k.SetAddressesCount(ctx, genState.AddressesCount)
	// Set all the tariff
	for _, elem := range genState.TariffList {
		k.SetTariff(ctx, elem)
	}
	// Set all the tariffs
	for _, elem := range genState.TariffsList {
		k.SetTariffs(ctx, elem)
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

	genesis.AddressesList = k.GetAllAddresses(ctx)
	genesis.AddressesCount = k.GetAddressesCount(ctx)
	genesis.TariffList = k.GetAllTariff(ctx)
	genesis.TariffsList = k.GetAllTariffs(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
