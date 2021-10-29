package keeper

import (
	"context"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) BoardComment(c context.Context, req *types.QueryGetBoardCommentRequest) (*types.QueryGetBoardCommentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetBoardComment(ctx, req.Id)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetBoardCommentResponse{BoardComment: &val}, nil
}
