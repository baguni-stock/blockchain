package types

import (
	"fmt"
	// this line is used by starport scaffolding # ibc/genesistype/import
)

// DefaultDate is the default capability global name
const DefaultDate uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		BoardCommentList:     []*BoardComment{},
		BoardList:            []*Board{},
		StockTransactionList: []*StockTransaction{},
		StockDataList:        []*StockData{},
		UserList:             []*User{},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated index in boardComment
	boardCommentIdMap := make(map[uint64]bool)

	for _, elem := range gs.BoardCommentList {
		if _, ok := boardCommentIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for boardComment")
		}
		boardCommentIdMap[elem.Id] = true
	}
	// Check for duplicated ID in board
	boardIdMap := make(map[uint64]bool)

	for _, elem := range gs.BoardList {
		if _, ok := boardIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for board")
		}
		boardIdMap[elem.Id] = true
	}
	// Check for duplicated index in stockTransaction
	stockTransactionIndexMap := make(map[string]bool)

	for _, elem := range gs.StockTransactionList {
		if _, ok := stockTransactionIndexMap[elem.Creator]; ok {
			return fmt.Errorf("duplicated Creaotr for stockTransaction")
		}
		stockTransactionIndexMap[elem.Creator] = true
	}
	// Check for duplicated index in stockData
	stockDataDateMap := make(map[string]bool)

	for _, elem := range gs.StockDataList {
		if _, ok := stockDataDateMap[elem.Date]; ok {
			return fmt.Errorf("duplicated index for stockData")
		}
		stockDataDateMap[elem.Date] = true
	}
	// Check for duplicated name in user
	userDateMap := make(map[string]bool)

	for _, elem := range gs.UserList {
		if _, ok := userDateMap[elem.Name]; ok {
			return fmt.Errorf("duplicated name for user")
		}
		userDateMap[elem.Name] = true
	}

	return nil
}
