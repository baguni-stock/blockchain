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
