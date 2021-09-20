package keeper

import (
	"context"
	"fmt"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateUser(goCtx context.Context, msg *types.MsgCreateUser) (*types.MsgCreateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetUser(ctx, msg.Name)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("name %v already set", msg.Name))
	}

	// check if creator already create user
	user_list := k.GetAllUser(ctx)
	for i := 0; i < len(user_list); i++ {
		if user_list[i].Creator == msg.Creator {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("address %v already registered", msg.Creator))
		}
	}

	var user = types.User{
		Name:    msg.Name,
		Creator: msg.Creator,
	}

	// coin 발행
	mint_coin, err := sdk.ParseCoinsNormalized("1000000token")
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Coin normalized faild")
	}
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mint_coin...))

	// creator address 얻기
	user_address, err := sdk.AccAddressFromBech32(user.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintln(err))
	}

	// creator에 coin제공
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user_address, mint_coin)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetUser(
		ctx,
		user,
	)
	return &types.MsgCreateUserResponse{}, nil
}

func (k msgServer) UpdateUser(goCtx context.Context, msg *types.MsgUpdateUser) (*types.MsgUpdateUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUser(ctx, msg.Name)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("date %v not set", msg.Name))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var user = types.User{
		Name:    msg.Name,
		Creator: msg.Creator,
	}

	k.SetUser(ctx, user)

	return &types.MsgUpdateUserResponse{}, nil
}

func (k msgServer) DeleteUser(goCtx context.Context, msg *types.MsgDeleteUser) (*types.MsgDeleteUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetUser(ctx, msg.Name)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("date %v not set", msg.Name))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveUser(ctx, msg.Name)

	return &types.MsgDeleteUserResponse{}, nil
}
