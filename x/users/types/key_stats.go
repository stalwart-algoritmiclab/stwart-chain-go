package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StatsKeyPrefix is the prefix to retrieve all Stats
	StatsKeyPrefix = "Stats/value/"
)

// StatsKey returns the store key to retrieve the Stats from the index fields
func StatsKey(
	date string,
) []byte {
	var key []byte

	dateBytes := []byte(date)
	key = append(key, dateBytes...)
	key = append(key, []byte("/")...)

	return key
}
