/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package types

const (
	// ModuleName defines the module name
	ModuleName = "stwart"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stwart"
)

var (
	ParamsKey = []byte("p_stwart")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
