/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		AddressesList: []Address{},
		TariffList:    []Tariff{},
		TariffsList:   []Tariffs{},
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
	// Check for duplicated index in tariff
	tariffIndexMap := make(map[string]struct{})

	for _, elem := range gs.TariffList {
		index := string(TariffKey(elem.Denom))
		if _, ok := tariffIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tariff")
		}
		tariffIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in tariffs
	tariffsIndexMap := make(map[string]struct{})

	for _, elem := range gs.TariffsList {
		index := string(TariffsKey(elem.Denom))
		if _, ok := tariffsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for tariffs")
		}
		tariffsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
