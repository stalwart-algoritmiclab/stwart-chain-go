package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TotalUsersKeyPrefix is the prefix to retrieve all TotalUsers
	TotalUsersKeyPrefix = "TotalUsers/value/"
)

// TotalUsersKey returns the store key to retrieve a TotalUsersKey from the index fields
func TotalUsersKey(
	date string,
) []byte {
	var key []byte

	dateBytes := []byte(date)
	key = append(key, dateBytes...)
	key = append(key, []byte("/")...)

	return key
}
