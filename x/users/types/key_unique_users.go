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
