package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// FeeStatsKeyPrefix is the prefix to retrieve all FeeStats
	FeeStatsKeyPrefix = "FeeStats/value/"
)

// FeeStatsKey returns the store key to retrieve a FeeStats from the index fields
func FeeStatsKey(
	date string,
) []byte {
	var key []byte

	dateBytes := []byte(date)
	key = append(key, dateBytes...)
	key = append(key, []byte("/")...)

	return key
}
