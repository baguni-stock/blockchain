package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetBoardCount get the total number of TypeName.LowerCamel
func (k Keeper) GetBoardCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardCountKey))
	byteKey := types.KeyPrefix(types.BoardCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to uint64
		panic("cannot decode count")
	}

	return count
}

// SetBoardCount set the total number of board
func (k Keeper) SetBoardCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardCountKey))
	byteKey := types.KeyPrefix(types.BoardCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendBoard appends a board in the store with a new id and update the count
func (k Keeper) AppendBoard(
	ctx sdk.Context,
	board types.Board,
) uint64 {
	// Create the board
	count := k.GetBoardCount(ctx)

	// Set the ID of the appended value
	board.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	appendedValue := k.cdc.MustMarshalBinaryBare(&board)
	store.Set(GetBoardIDBytes(board.Id), appendedValue)

	// Update board count
	k.SetBoardCount(ctx, count+1)

	return count
}

// SetBoard set a specific board in the store
func (k Keeper) SetBoard(ctx sdk.Context, board types.Board) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	b := k.cdc.MustMarshalBinaryBare(&board)
	store.Set(GetBoardIDBytes(board.Id), b)
}

// GetBoard returns a board from its id
func (k Keeper) GetBoard(ctx sdk.Context, id uint64) types.Board {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	var board types.Board
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBoardIDBytes(id)), &board)
	return board
}

// HasBoard checks if the board exists in the store
func (k Keeper) HasBoard(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	return store.Has(GetBoardIDBytes(id))
}

// GetBoardOwner returns the creator of the
func (k Keeper) GetBoardOwner(ctx sdk.Context, id uint64) string {
	return k.GetBoard(ctx, id).Creator
}

// RemoveBoard removes a board from the store
func (k Keeper) RemoveBoard(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	store.Delete(GetBoardIDBytes(id))
}

// GetAllBoard returns all board
func (k Keeper) GetAllBoard(ctx sdk.Context) (list []types.Board) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BoardKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Board
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBoardIDBytes returns the byte representation of the ID
func GetBoardIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBoardIDFromBytes returns ID in uint64 format from a byte array
func GetBoardIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
