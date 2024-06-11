/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UserKeyPrefix is the prefix to retrieve all User
	UserKeyPrefix = "User/value/"
)

// UserKey returns the store key to retrieve a User from the index fields
func UserKey(
	accountAddress string,
) []byte {
	var key []byte

	accountAddressBytes := []byte(accountAddress)
	key = append(key, accountAddressBytes...)
	key = append(key, []byte("/")...)

	return key
}
