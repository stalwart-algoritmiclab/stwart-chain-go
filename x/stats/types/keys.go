package types

const (
	// ModuleName defines the module name
	ModuleName = "stats"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stats"
)

var (
	ParamsKey = []byte("p_stats")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
