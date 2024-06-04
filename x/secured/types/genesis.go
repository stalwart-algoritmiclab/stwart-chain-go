package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AddressesList: []Addresses{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in addresses
	addressesIdMap := make(map[uint64]bool)
	addressesCount := gs.GetAddressesCount()
	for _, elem := range gs.AddressesList {
		if _, ok := addressesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for addresses")
		}
		if elem.Id >= addressesCount {
			return fmt.Errorf("addresses id should be lower or equal than the last id")
		}
		addressesIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
