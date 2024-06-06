/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
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
		StatsList:       []Stats{},
		UniqueUsersList: []UniqueUsers{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in stats
	statsIndexMap := make(map[string]struct{})

	for _, elem := range gs.StatsList {
		index := string(StatsKey(elem.Date))
		if _, ok := statsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for stats")
		}
		statsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in uniqueUsers
	uniqueUsersIndexMap := make(map[string]struct{})

	for _, elem := range gs.UniqueUsersList {
		index := string(UniqueUsersKey(elem.Date))
		if _, ok := uniqueUsersIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for uniqueUsers")
		}
		uniqueUsersIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
