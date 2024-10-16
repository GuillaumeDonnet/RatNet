package types

const (
	// ModuleName defines the module name
	ModuleName = "labrat"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_labrat"
)

var (
	ParamsKey = []byte("p_labrat")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
