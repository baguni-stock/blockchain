package types

const (
	// ModuleName defines the module name
	ModuleName = "blockchain"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_blockchain"

	// this line is used by starport scaffolding # ibc/keys/name
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	UserKey = "User-value-"
)

const (
	StockDataKey = "StockData-value-"
)

const (
	StockTransactionKey = "StockTransaction-value-"
)
