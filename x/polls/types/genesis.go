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
		PollsParams: nil,
		VotesList:   []Votes{},
		OptionsList: []Options{},
		PollsList:   []Polls{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in votes
	votesIdMap := make(map[uint64]bool)
	votesCount := gs.GetVotesCount()
	for _, elem := range gs.VotesList {
		if _, ok := votesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for votes")
		}
		if elem.Id >= votesCount {
			return fmt.Errorf("votes id should be lower or equal than the last id")
		}
		votesIdMap[elem.Id] = true
	}
	// Check for duplicated ID in options
	optionsIdMap := make(map[uint64]bool)
	optionsCount := gs.GetOptionsCount()
	for _, elem := range gs.OptionsList {
		if _, ok := optionsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for options")
		}
		if elem.Id >= optionsCount {
			return fmt.Errorf("options id should be lower or equal than the last id")
		}
		optionsIdMap[elem.Id] = true
	}
	// Check for duplicated ID in polls
	pollsIdMap := make(map[uint64]bool)
	pollsCount := gs.GetPollsCount()
	for _, elem := range gs.PollsList {
		if _, ok := pollsIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for polls")
		}
		if elem.Id >= pollsCount {
			return fmt.Errorf("polls id should be lower or equal than the last id")
		}
		pollsIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
