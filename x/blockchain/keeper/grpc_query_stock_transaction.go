package keeper

import (
	"context"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StockTransactionAll(c context.Context, req *types.QueryAllStockTransactionRequest) (*types.QueryAllStockTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stockTransactions []*types.StockTransaction
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	stockTransactionStore := prefix.NewStore(store, types.KeyPrefix(types.StockTransactionKey))

	pageRes, err := query.Paginate(stockTransactionStore, req.Pagination, func(key []byte, value []byte) error {
		var stockTransaction types.StockTransaction
		if err := k.cdc.UnmarshalBinaryBare(value, &stockTransaction); err != nil {
			return err
		}

		stockTransactions = append(stockTransactions, &stockTransaction)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStockTransactionResponse{StockTransaction: stockTransactions, Pagination: pageRes}, nil
}

func (k Keeper) StockTransaction(c context.Context, req *types.QueryGetStockTransactionRequest) (*types.QueryGetStockTransactionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStockTransaction(ctx, req.Creator)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetStockTransactionResponse{StockTransaction: &val}, nil
}
