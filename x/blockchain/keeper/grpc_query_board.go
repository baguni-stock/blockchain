package keeper

import (
	"context"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BoardAll(c context.Context, req *types.QueryAllBoardRequest) (*types.QueryAllBoardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var boards []*types.Board
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	boardStore := prefix.NewStore(store, types.KeyPrefix(types.BoardKey))

	pageRes, err := query.Paginate(boardStore, req.Pagination, func(key []byte, value []byte) error {
		var board types.Board
		if err := k.cdc.UnmarshalBinaryBare(value, &board); err != nil {
			return err
		}

		boards = append(boards, &board)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllBoardResponse{Board: boards, Pagination: pageRes}, nil
}

func (k Keeper) Board(c context.Context, req *types.QueryGetBoardRequest) (*types.QueryGetBoardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var board types.Board
	ctx := sdk.UnwrapSDKContext(c)

	if !k.HasBoard(ctx, req.Id) {
		return nil, sdkerrors.ErrKeyNotFound
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBoardIDBytes(req.Id)), &board)

	return &types.QueryGetBoardResponse{Board: &board}, nil
}
