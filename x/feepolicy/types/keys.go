package types

const (
	// ModuleName defines the module name
	ModuleName = "feepolicy"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_feepolicy"
)

var (
	ParamsKey = []byte("p_feepolicy")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	AddressesKey      = "Addresses/value/"
	AddressesCountKey = "Addresses/count/"
)
