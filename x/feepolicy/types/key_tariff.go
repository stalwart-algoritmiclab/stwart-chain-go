package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TariffKeyPrefix is the prefix to retrieve all Tariff
	TariffKeyPrefix = "Tariff/value/"
)

// TariffKey returns the store key to retrieve a Tariff from the index fields
func TariffKey(
	denom string,
) []byte {
	var key []byte

	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
