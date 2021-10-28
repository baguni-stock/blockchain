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
func (k Keeper) GetStockTransaction(ctx sdk.Context, creator string) (val types.StockTransaction, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockTransactionKey))

	b := store.Get(types.KeyPrefix(creator))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveStockTransaction removes a stockTransaction from the store
func (k Keeper) RemoveStockTransaction(ctx sdk.Context, creator string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockTransactionKey))
	store.Delete(types.KeyPrefix(creator))
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

func (k Keeper) SetStockTransactionRecord(ctx sdk.Context, stockTransactionRecord types.StockTransactionRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockTransactionRecordKey))
	b := k.cdc.MustMarshalBinaryBare(&stockTransactionRecord)
	store.Set(types.KeyPrefix(stockTransactionRecord.Creator), b)
}

// GetStockTransaction returns a stockTransaction from its creator
func (k Keeper) GetStockTransactionRecord(ctx sdk.Context, creator string) (val types.StockTransactionRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockTransactionRecordKey))

	b := store.Get(types.KeyPrefix(creator))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// GetAllStockTransaction returns all stockTransaction
func (k Keeper) GetAllStockTransactionRecord(ctx sdk.Context) (list []types.StockTransactionRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StockTransactionRecordKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StockTransactionRecord
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
