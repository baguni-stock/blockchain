package keeper

import (
	"context"
	"fmt"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateBoard(goCtx context.Context, msg *types.MsgCreateBoard) (*types.MsgCreateBoardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var board = types.Board{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	id := k.AppendBoard(
		ctx,
		board,
	)

	// coments 등록
	var comments []*types.Comment
	board_comment := types.BoardComment{
		Creator: msg.Creator,
		Id:      id,
		Count:   0,
		Coments: comments,
	}
	k.SetBoardComment(ctx, board_comment)

	return &types.MsgCreateBoardResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateBoard(goCtx context.Context, msg *types.MsgUpdateBoard) (*types.MsgUpdateBoardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var board = types.Board{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	// Checks that the element exists
	if !k.HasBoard(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != k.GetBoardOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetBoard(ctx, board)

	return &types.MsgUpdateBoardResponse{}, nil
}

func (k msgServer) DeleteBoard(goCtx context.Context, msg *types.MsgDeleteBoard) (*types.MsgDeleteBoardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasBoard(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != k.GetBoardOwner(ctx, msg.Id) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveBoard(ctx, msg.Id)
	k.RemoveBoardComment(ctx, msg.Id)
	return &types.MsgDeleteBoardResponse{}, nil
}
