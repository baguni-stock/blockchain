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

	root_address := "cosmos1s3pzgpduvnq4r59mjx0vmdzfttqkhywwj7f8lk"
	if root_address != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("creator not root %+v", msg.Creator))
	}

	// Check if the value already exists_
	_, isFound := k.GetUser(ctx, msg.Name)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("name %v already set", msg.Name))
	}

	// check if creator already create user
	user_list := k.GetAllUser(ctx)
	for i := 0; i < len(user_list); i++ {
		if user_list[i].Address == msg.Address {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("address %v already registered", msg.Address))
		}
	}

	var user = types.User{
		Creator: msg.Creator,
		Address: msg.Address,
		Name:    msg.Name,
	}

	// coin 발행
	mint_coin, err := sdk.ParseCoinsNormalized("1000000stake")
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Coin normalized faild")
	}
	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mint_coin...))

	// address to account
	user_address, err := sdk.AccAddressFromBech32(user.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintln(err))
	}

	// address에 coin제공
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
		Creator: msg.Creator,
		Address: msg.Address,
		Name:    msg.Name,
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
