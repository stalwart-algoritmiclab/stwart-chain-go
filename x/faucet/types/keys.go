/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package types

const (
	// ModuleName defines the module name
	ModuleName = "faucet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_faucet"
)

var (
	ParamsKey = []byte("p_faucet")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	TokensKey      = "Tokens/value/"
	TokensCountKey = "Tokens/count/"
)
