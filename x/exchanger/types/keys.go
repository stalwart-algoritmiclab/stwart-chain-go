package types

const (
	// ModuleName defines the module name
	ModuleName = "exchanger"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_exchanger"
)

var (
	ParamsKey = []byte("p_exchanger")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
