package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TariffsKeyPrefix is the prefix to retrieve all Tariffs
	TariffsKeyPrefix = "Tariffs/value/"
)

// TariffsKey returns the store key to retrieve a Tariffs from the index fields
func TariffsKey(
	denom string,
) []byte {
	var key []byte

	denomBytes := []byte(denom)
	key = append(key, denomBytes...)
	key = append(key, []byte("/")...)

	return key
}
