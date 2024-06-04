package types

const (
	// ModuleName defines the module name
	ModuleName = "core"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_core"
)

var (
	ParamsKey = []byte("p_core")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
