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

func (k Keeper) StockDataAll(c context.Context, req *types.QueryAllStockDataRequest) (*types.QueryAllStockDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var stockDatas []*types.StockData
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	stockDataStore := prefix.NewStore(store, types.KeyPrefix(types.StockDataKey))

	pageRes, err := query.Paginate(stockDataStore, req.Pagination, func(key []byte, value []byte) error {
		var stockData types.StockData
		if err := k.cdc.UnmarshalBinaryBare(value, &stockData); err != nil {
			return err
		}

		stockDatas = append(stockDatas, &stockData)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStockDataResponse{StockData: stockDatas, Pagination: pageRes}, nil
}

func (k Keeper) StockData(c context.Context, req *types.QueryGetStockDataRequest) (*types.QueryGetStockDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetStockData(ctx, req.Date)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetStockDataResponse{StockData: &val}, nil
}
