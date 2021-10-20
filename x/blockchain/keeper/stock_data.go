package keeper

import (
	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStockData set a specific stockData in the store from its date
func (k Keeper) SetStockData(ctx sdk.Context, stockData types.StockData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockDataKey))
	b := k.cdc.MustMarshalBinaryBare(&stockData)
	store.Set(types.KeyPrefix(stockData.Code), b)
}

// GetStockData returns a stockData from its date
func (k Keeper) GetStockData(ctx sdk.Context, date string) (val types.StockData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockDataKey))

	b := store.Get(types.KeyPrefix(date))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// GetStockData returns a stockData from its date
func (k Keeper) GetLastStockData(ctx sdk.Context, date string) (val types.StockData, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockDataKey))

	b := store.Get(types.KeyPrefix(date))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveStockData removes a stockData from the store
func (k Keeper) RemoveStockData(ctx sdk.Context, date string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockDataKey))
	store.Delete(types.KeyPrefix(date))
}

// GetAllStockData returns all stockData
func (k Keeper) GetAllStockData(ctx sdk.Context) (list []types.StockData) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockDataKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StockData
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
