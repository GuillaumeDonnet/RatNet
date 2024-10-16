package types

const (
	// ModuleName defines the module name
	ModuleName = "ratnet"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ratnet"
)

var (
	ParamsKey = []byte("p_ratnet")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
