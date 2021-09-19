package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateStockData{}

func NewMsgCreateStockData(creator string, index string, stocks []*Stock) *MsgCreateStockData {
	return &MsgCreateStockData{
		Creator: creator,
		Index:   index,
		Stocks:  stocks,
	}
}

func (msg *MsgCreateStockData) Route() string {
	return RouterKey
}

func (msg *MsgCreateStockData) Type() string {
	return "CreateStockData"
}

func (msg *MsgCreateStockData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateStockData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateStockData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateStockData{}

func NewMsgUpdateStockData(creator string, index string, stocks []*Stock) *MsgUpdateStockData {
	return &MsgUpdateStockData{
		Creator: creator,
		Index:   index,
		Stocks:  stocks,
	}
}

func (msg *MsgUpdateStockData) Route() string {
	return RouterKey
}

func (msg *MsgUpdateStockData) Type() string {
	return "UpdateStockData"
}

func (msg *MsgUpdateStockData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateStockData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateStockData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteStockData{}

func NewMsgDeleteStockData(creator string, index string) *MsgDeleteStockData {
	return &MsgDeleteStockData{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteStockData) Route() string {
	return RouterKey
}

func (msg *MsgDeleteStockData) Type() string {
	return "DeleteStockData"
}

func (msg *MsgDeleteStockData) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteStockData) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteStockData) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
