package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateBoard{}

func NewMsgCreateBoard(creator string, title string, body string) *MsgCreateBoard {
	return &MsgCreateBoard{
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgCreateBoard) Route() string {
	return RouterKey
}

func (msg *MsgCreateBoard) Type() string {
	return "CreateBoard"
}

func (msg *MsgCreateBoard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateBoard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateBoard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateBoard{}

func NewMsgUpdateBoard(creator string, id uint64, title string, body string) *MsgUpdateBoard {
	return &MsgUpdateBoard{
		Id:      id,
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgUpdateBoard) Route() string {
	return RouterKey
}

func (msg *MsgUpdateBoard) Type() string {
	return "UpdateBoard"
}

func (msg *MsgUpdateBoard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateBoard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateBoard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteBoard{}

func NewMsgDeleteBoard(creator string, id uint64) *MsgDeleteBoard {
	return &MsgDeleteBoard{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteBoard) Route() string {
	return RouterKey
}

func (msg *MsgDeleteBoard) Type() string {
	return "DeleteBoard"
}

func (msg *MsgDeleteBoard) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteBoard) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteBoard) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
