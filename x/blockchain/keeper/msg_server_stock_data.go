package keeper

import (
	"context"
	"fmt"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStockData(goCtx context.Context, msg *types.MsgCreateStockData) (*types.MsgCreateStockDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Croot address 인지 체크
	root_address := "cosmos1s3pzgpduvnq4r59mjx0vmdzfttqkhywwj7f8lk"
	if root_address != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("creator not root %+v", msg.Creator))
	}

	for i := 0; i < len(msg.Stocks); i++ {
		k.SetStockData(
			ctx,
			*msg.Stocks[i],
		)
	}
	return &types.MsgCreateStockDataResponse{}, nil
}

func (k msgServer) DeleteStockData(goCtx context.Context, msg *types.MsgDeleteStockData) (*types.MsgDeleteStockDataResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStockData(ctx, msg.Code)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("Code %v not set", msg.Code))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveStockData(ctx, msg.Code)
	return &types.MsgDeleteStockDataResponse{}, nil
}
