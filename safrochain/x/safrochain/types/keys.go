package types

const (
	// ModuleName defines the module name
	ModuleName = "safrochain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_safrochain"
)

var (
	ParamsKey = []byte("p_safrochain")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
