package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateStockTransaction{}

func NewMsgCreateStockTransaction(creator string, code int32, count int32) *MsgCreateStockTransaction {
	return &MsgCreateStockTransaction{
		Creator: creator,
		Code:    code,
		Count:   count,
	}
}

func (msg *MsgCreateStockTransaction) Route() string {
	return RouterKey
}

func (msg *MsgCreateStockTransaction) Type() string {
	return "CreateStockTransaction"
}

func (msg *MsgCreateStockTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateStockTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateStockTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteStockTransaction{}

func NewMsgDeleteStockTransaction(creator string, code int32, count int32) *MsgDeleteStockTransaction {
	return &MsgDeleteStockTransaction{
		Creator: creator,
		Code:    code,
		Count:   count,
	}
}
func (msg *MsgDeleteStockTransaction) Route() string {
	return RouterKey
}

func (msg *MsgDeleteStockTransaction) Type() string {
	return "DeleteStockTransaction"
}

func (msg *MsgDeleteStockTransaction) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteStockTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteStockTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
