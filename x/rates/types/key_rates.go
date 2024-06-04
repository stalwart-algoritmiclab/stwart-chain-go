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
