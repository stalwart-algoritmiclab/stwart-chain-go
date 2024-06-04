package types

const (
	// ModuleName defines the module name
	ModuleName = "referral"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_referral"
)

var (
	ParamsKey = []byte("p_referral")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
