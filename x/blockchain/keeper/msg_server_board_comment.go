package keeper

import (
	"context"
	"fmt"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateBoardComment(goCtx context.Context, msg *types.MsgCreateBoardComment) (*types.MsgCreateBoardCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetBoardComment(ctx, msg.BoardCommentId)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("board_comment_id %v board not set", msg.BoardCommentId))
	}

	var comment = types.Comment{
		Creator: msg.Creator,
		Body:    msg.Body,
	}
	k.AppendComment(ctx, msg.BoardCommentId, comment)

	return &types.MsgCreateBoardCommentResponse{}, nil
}

func (k msgServer) UpdateBoardComment(goCtx context.Context, msg *types.MsgUpdateBoardComment) (*types.MsgUpdateBoardCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetBoardComment(ctx, msg.BoardCommentId)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("board_comment_id %v not set", msg.BoardCommentId))
	}

	// Checks if the the msg sender is the same as the current owner
	comment := valFound.Coments[msg.CommentId]
	if msg.Creator != comment.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	comment = &types.Comment{
		Creator: msg.Creator,
		Id:      msg.CommentId,
		Body:    msg.Body,
	}
	k.SetComment(ctx, msg.BoardCommentId, comment)

	return &types.MsgUpdateBoardCommentResponse{}, nil
}

func (k msgServer) DeleteBoardComment(goCtx context.Context, msg *types.MsgDeleteBoardComment) (*types.MsgDeleteBoardCommentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetBoardComment(ctx, msg.BoardCommentId)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("board_comment_id %v not set", msg.BoardCommentId))
	}

	// Checks if the the msg sender is the same as the current owner
	comment := valFound.Coments[msg.CommentId]
	if msg.Creator != comment.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveComment(ctx, msg.BoardCommentId, msg.CommentId)

	return &types.MsgDeleteBoardCommentResponse{}, nil
}
