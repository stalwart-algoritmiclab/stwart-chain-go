/*
 * SPDX-License-Identifier: BUSL-1.1
 * Contributed by Algoritmic Lab Ltd. Copyright (C) 2024.
 * Full license is available at https://github.com/stalwart-algoritmiclab/stwart-chain-go/blob/main/LICENCE
 */

package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// RatesKeyPrefix is the prefix to retrieve all Rates
	RatesKeyPrefix = "Rates/value/"
)

// RatesKey returns the store key to retrieve a Rates from the index fields
func RatesKey(
	denom string,
) []byte {
	var key []byte

	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
