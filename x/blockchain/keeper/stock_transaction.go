package keeper

import (
	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStockTransaction set a specific stockTransaction in the store from its index
func (k Keeper) SetStockTransaction(ctx sdk.Context, stockTransaction types.StockTransaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockTransactionKey))
	b := k.cdc.MustMarshalBinaryBare(&stockTransaction)
	store.Set(types.KeyPrefix(stockTransaction.Creator), b)
}

// GetStockTransaction returns a stockTransaction from its index
func (k Keeper) GetStockTransaction(ctx sdk.Context, index string) (val types.StockTransaction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockTransactionKey))

	b := store.Get(types.KeyPrefix(index))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveStockTransaction removes a stockTransaction from the store
func (k Keeper) RemoveStockTransaction(ctx sdk.Context, index string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockTransactionKey))
	store.Delete(types.KeyPrefix(index))
}

// GetAllStockTransaction returns all stockTransaction
func (k Keeper) GetAllStockTransaction(ctx sdk.Context) (list []types.StockTransaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockTransactionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StockTransaction
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
