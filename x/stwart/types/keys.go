package types

const (
	// ModuleName defines the module name
	ModuleName = "stwart"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stwart"
)

var (
	ParamsKey = []byte("p_stwart")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
