package keeper

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
)

func createNStockData(keeper *Keeper, ctx sdk.Context, n int) []types.StockData {
	items := make([]types.StockData, n)
	for i := range items {
		items[i].Creator = "any"
		items[i].Index = fmt.Sprintf("%d", i)
		keeper.SetStockData(ctx, items[i])
	}
	return items
}

func TestStockDataGet(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNStockData(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStockData(ctx, item.Index)
		assert.True(t, found)
		assert.Equal(t, item, rst)
	}
}
func TestStockDataRemove(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNStockData(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStockData(ctx, item.Index)
		_, found := keeper.GetStockData(ctx, item.Index)
		assert.False(t, found)
	}
}

func TestStockDataGetAll(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	items := createNStockData(keeper, ctx, 10)
	assert.Equal(t, items, keeper.GetAllStockData(ctx))
}
