package blockchain

import (
	"github.com/chainstock-project/blockchain/x/blockchain/keeper"
	"github.com/chainstock-project/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	// Set all the boardComment
	for _, elem := range genState.BoardCommentList {
		k.SetBoardComment(ctx, *elem)
	}

	// Set all the board
	for _, elem := range genState.BoardList {
		k.SetBoard(ctx, *elem)
	}

	// Set board count
	k.SetBoardCount(ctx, genState.BoardCount)

	// Set all the stockTransaction
	for _, elem := range genState.StockTransactionList {
		k.SetStockTransaction(ctx, *elem)
	}

	// Set all the stockData
	for _, elem := range genState.StockDataList {
		k.SetStockData(ctx, *elem)
	}

	// Set all the user
	for _, elem := range genState.UserList {
		k.SetUser(ctx, *elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export
	// Get all boardComment
	boardCommentList := k.GetAllBoardComment(ctx)
	for _, elem := range boardCommentList {
		elem := elem
		genesis.BoardCommentList = append(genesis.BoardCommentList, &elem)
	}

	// Get all board
	boardList := k.GetAllBoard(ctx)
	for _, elem := range boardList {
		elem := elem
		genesis.BoardList = append(genesis.BoardList, &elem)
	}

	// Set the current count
	genesis.BoardCount = k.GetBoardCount(ctx)

	// Get all stockTransaction
	stockTransactionList := k.GetAllStockTransaction(ctx)
	for _, elem := range stockTransactionList {
		elem := elem
		genesis.StockTransactionList = append(genesis.StockTransactionList, &elem)
	}

	// Get all stockData
	stockDataList := k.GetAllStockData(ctx)
	for _, elem := range stockDataList {
		elem := elem
		genesis.StockDataList = append(genesis.StockDataList, &elem)
	}

	// Get all user
	userList := k.GetAllUser(ctx)
	for _, elem := range userList {
		elem := elem
		genesis.UserList = append(genesis.UserList, &elem)
	}

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
