package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBoardComment{}

func NewMsgCreateBoardComment(creator string, board_comment_id uint64, body string) *MsgCreateBoardComment {
	return &MsgCreateBoardComment{
		Creator:        creator,
		BoardCommentId: board_comment_id,
		Body:           body,
	}
}

func (msg *MsgCreateBoardComment) Route() string {
	return RouterKey
}

func (msg *MsgCreateBoardComment) Type() string {
	return "CreateBoardComment"
}

func (msg *MsgCreateBoardComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBoardComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBoardComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBoardComment{}

func NewMsgUpdateBoardComment(creator string, board_comment_id uint64, comment_id int64, body string) *MsgUpdateBoardComment {
	return &MsgUpdateBoardComment{
		Creator:        creator,
		BoardCommentId: board_comment_id,
		CommentId:      comment_id,
		Body:           body,
	}
}

func (msg *MsgUpdateBoardComment) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBoardComment) Type() string {
	return "UpdateBoardComment"
}

func (msg *MsgUpdateBoardComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBoardComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBoardComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBoardComment{}

func NewMsgDeleteBoardComment(creator string, board_comment_id uint64, comment_id int64) *MsgDeleteBoardComment {
	return &MsgDeleteBoardComment{
		Creator:        creator,
		BoardCommentId: board_comment_id,
		CommentId:      comment_id,
	}
}
func (msg *MsgDeleteBoardComment) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBoardComment) Type() string {
	return "DeleteBoardComment"
}

func (msg *MsgDeleteBoardComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBoardComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBoardComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
