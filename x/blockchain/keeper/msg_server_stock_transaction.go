package keeper

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStockTransaction(goCtx context.Context, msg *types.MsgCreateStockTransaction) (*types.MsgCreateStockTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	stock_transaction, isFound := k.GetStockTransaction(ctx, msg.Creator)
	if !isFound {
		var holding_stocks []*types.HoldingStock
		stock_transaction = types.StockTransaction{
			Creator:       msg.Creator,
			HoldingStocks: holding_stocks,
		}
	}

	//stock code의 가격(amount) 얻기
	var amount string = ""

	utc := time.Now().UTC()
	loc, _ := time.LoadLocation("Asia/Seoul")
	kst := utc.In(loc)
	date := kst.Format("2006-01-02")
	stock_data, isFound := k.GetStockData(ctx, date)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "incrrect date")

	}

	for i := 0; i < len(stock_data.Stocks); i++ {
		if stock_data.Stocks[i].Code == msg.Code {
			amount = stock_data.Stocks[i].Amount
			break
		}
	}
	if amount == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "incorrect amount")
	}

	// 주식 coin 지불
	count, _ := strconv.Atoi(msg.Count)
	stake, _ := strconv.Atoi(amount)
	stake *= count
	stake_string := strconv.Itoa(stake)

	burn_coin, err := sdk.ParseCoinsNormalized(stake_string + "stake")
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Coin normalized faild")
	}

	user_address, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintln(err))
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, user_address, types.ModuleName, burn_coin)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintln(err))
	}
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, burn_coin)

	holding_stock := types.HoldingStock{
		Code:   msg.Code,
		Count:  msg.Count,
		Amount: amount,
	}
	stock_transaction.HoldingStocks = append(stock_transaction.HoldingStocks, &holding_stock)

	k.SetStockTransaction(
		ctx,
		stock_transaction,
	)
	return &types.MsgCreateStockTransactionResponse{}, nil

}

func (k msgServer) DeleteStockTransaction(goCtx context.Context, msg *types.MsgDeleteStockTransaction) (*types.MsgDeleteStockTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetStockTransaction(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator %v not set", msg.Creator))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveStockTransaction(ctx, msg.Creator)

	return &types.MsgDeleteStockTransactionResponse{}, nil
}
