/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

const (
	// ModuleName defines the module name
	ModuleName = "stake"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stake"
)

var (
	ParamsKey = []byte("p_stake")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
