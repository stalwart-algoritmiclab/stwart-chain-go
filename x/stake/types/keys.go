package types

const (
	// ModuleName defines the module name
	ModuleName = "stake"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stake"
)

var (
	ParamsKey = []byte("p_stake")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
