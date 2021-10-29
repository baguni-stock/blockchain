package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateBoardComment{}, "blockchain/CreateBoardComment", nil)
	cdc.RegisterConcrete(&MsgUpdateBoardComment{}, "blockchain/UpdateBoardComment", nil)
	cdc.RegisterConcrete(&MsgDeleteBoardComment{}, "blockchain/DeleteBoardComment", nil)

	cdc.RegisterConcrete(&MsgCreateBoard{}, "blockchain/CreateBoard", nil)
	cdc.RegisterConcrete(&MsgUpdateBoard{}, "blockchain/UpdateBoard", nil)
	cdc.RegisterConcrete(&MsgDeleteBoard{}, "blockchain/DeleteBoard", nil)

	cdc.RegisterConcrete(&MsgCreateStockTransaction{}, "blockchain/CreateStockTransaction", nil)
	cdc.RegisterConcrete(&MsgDeleteStockTransaction{}, "blockchain/DeleteStockTransaction", nil)

	cdc.RegisterConcrete(&MsgCreateStockData{}, "blockchain/CreateStockData", nil)
	cdc.RegisterConcrete(&MsgDeleteStockData{}, "blockchain/DeleteStockData", nil)

	cdc.RegisterConcrete(&MsgCreateUser{}, "blockchain/CreateUser", nil)
	cdc.RegisterConcrete(&MsgUpdateUser{}, "blockchain/UpdateUser", nil)
	cdc.RegisterConcrete(&MsgDeleteUser{}, "blockchain/DeleteUser", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBoardComment{},
		&MsgUpdateBoardComment{},
		&MsgDeleteBoardComment{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateBoard{},
		&MsgUpdateBoard{},
		&MsgDeleteBoard{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateStockTransaction{},
		&MsgDeleteStockTransaction{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateStockData{},
		&MsgDeleteStockData{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUser{},
		&MsgUpdateUser{},
		&MsgDeleteUser{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
