package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		FeeStatsList: []FeeStats{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in feeStats
	feeStatsIndexMap := make(map[string]struct{})

	for _, elem := range gs.FeeStatsList {
		index := string(FeeStatsKey(elem.Date))
		if _, ok := feeStatsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for feeStats")
		}
		feeStatsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}