/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

const (
	// ModuleName defines the module name
	ModuleName = "polls"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_polls"
)

var (
	ParamsKey = []byte("p_polls")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PollsParamsKey = "PollsParams/value/"
)

const (
	VotesKey      = "Votes/value/"
	VotesCountKey = "Votes/count/"
)

const (
	OptionsKey      = "Options/value/"
	OptionsCountKey = "Options/count/"
)

const (
	PollsKey      = "Polls/value/"
	PollsCountKey = "Polls/count/"
)
