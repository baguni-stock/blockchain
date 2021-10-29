package keeper

import (
	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetBoardComment set a specific boardComment in the store from its id
func (k Keeper) SetBoardComment(ctx sdk.Context, boardComment types.BoardComment) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardCommentKey))
	b := k.cdc.MustMarshalBinaryBare(&boardComment)
	store.Set(GetBoardIDBytes(boardComment.Id), b)
}

// GetBoardComment returns a boardComment from its id
func (k Keeper) GetBoardComment(ctx sdk.Context, id uint64) (val types.BoardComment, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardCommentKey))

	b := store.Get(GetBoardIDBytes(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshalBinaryBare(b, &val)
	return val, true
}

// RemoveBoardComment removes a boardComment from the store
func (k Keeper) RemoveBoardComment(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardCommentKey))
	store.Delete(GetBoardIDBytes(id))
}

// GetAllBoardComment returns all boardComment
func (k Keeper) GetAllBoardComment(ctx sdk.Context) (list []types.BoardComment) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardCommentKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BoardComment
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) GetCommentCount(ctx sdk.Context, board_comment_id uint64) int64 {
	board_comment, isFound := k.GetBoardComment(ctx, board_comment_id)
	// Count doesn't exist: no element
	if !isFound {
		return 0
	}

	return board_comment.Count
}

// SetBoardCount set the total number of board
func (k Keeper) SetCommentCount(ctx sdk.Context, board_comment_id uint64, count int64) {
	board_comment, isFound := k.GetBoardComment(ctx, board_comment_id)
	// Count doesn't exist: no element
	if !isFound {
		return
	}

	board_comment.Count = count
	k.SetBoardComment(ctx, board_comment)
}

func (k Keeper) AppendComment(ctx sdk.Context, board_comment_id uint64, comment types.Comment) {
	count := k.GetCommentCount(ctx, board_comment_id)
	// Set the ID of the appended value
	comment.Id = count

	boardComment, _ := k.GetBoardComment(ctx, board_comment_id)
	boardComment.Coments = append(boardComment.Coments, &comment)
	k.SetBoardComment(ctx, boardComment)

	// Update comment count
	k.SetCommentCount(ctx, board_comment_id, count+1)
}

func (k Keeper) SetComment(ctx sdk.Context, board_comment_id uint64, comment *types.Comment) {
	boardComment, _ := k.GetBoardComment(ctx, board_comment_id)
	if boardComment.Coments[comment.Id].Id == -1 {
		panic("deleted")
	}
	boardComment.Coments[comment.Id] = comment
	k.SetBoardComment(
		ctx,
		boardComment,
	)
}

// GetBoardComment returns a boardComment from its id
func (k Keeper) GetComment(ctx sdk.Context, board_comment_id uint64, comment_id int64) (val *types.Comment, found bool) {
	boardComment, _ := k.GetBoardComment(ctx, board_comment_id)
	val = boardComment.Coments[comment_id]
	if val == nil || val.Id == -1 {
		return nil, false
	}
	return val, true
}

// RemoveBoardComment removes a boardComment from the store
func (k Keeper) RemoveComment(ctx sdk.Context, board_comment_id uint64, comment_id int64) {
	boardComment, _ := k.GetBoardComment(ctx, board_comment_id)
	boardComment.Coments[comment_id].Id = -1
	k.SetBoardComment(ctx, boardComment)
}
