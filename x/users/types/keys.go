package types

const (
	// ModuleName defines the module name
	ModuleName = "users"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_users"
)

var (
	ParamsKey = []byte("p_users")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
