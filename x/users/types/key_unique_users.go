/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/tree/main/LICENSES
 */

package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UniqueUsersKeyPrefix is the prefix to retrieve all UniqueUsers
	UniqueUsersKeyPrefix = "UniqueUsers/value/"
)

// UniqueUsersKey returns the store key to retrieve a UniqueUsers from the index fields
func UniqueUsersKey(
	date string,
) []byte {
	var key []byte

	dateBytes := []byte(date)
	key = append(key, dateBytes...)
	key = append(key, []byte("/")...)

	return key
}
