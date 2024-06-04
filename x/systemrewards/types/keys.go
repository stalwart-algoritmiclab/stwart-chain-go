package types

const (
	// ModuleName defines the module name
	ModuleName = "systemrewards"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_systemrewards"
)

var (
	ParamsKey = []byte("p_systemrewards")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
