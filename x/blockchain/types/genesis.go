package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultIndex is the default capability global name
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		StockDataList: []*StockData{},
		UserList:      []*User{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in stockData
	stockDataIndexMap := make(map[string]bool)

	for _, elem := range gs.StockDataList {
		if _, ok := stockDataIndexMap[elem.Index]; ok {
			return fmt.Errorf("duplicated index for stockData")
		}
		stockDataIndexMap[elem.Index] = true
	}
	// Check for duplicated name in user
	userIndexMap := make(map[string]bool)

	for _, elem := range gs.UserList {
		if _, ok := userIndexMap[elem.Name]; ok {
			return fmt.Errorf("duplicated name for user")
		}
		userIndexMap[elem.Name] = true
	}

	return nil
}
