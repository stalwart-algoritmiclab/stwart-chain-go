/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StakeKeyPrefix is the prefix to retrieve all Stake
	StakeKeyPrefix = "Stake/value/"
)

// StakeKey returns the store key to retrieve a Stake from the index fields
func StakeKey(
	address string,
) []byte {
	var key []byte

	addressBytes := []byte(address)
	key = append(key, addressBytes...)
	key = append(key, []byte("/")...)

	return key
}
