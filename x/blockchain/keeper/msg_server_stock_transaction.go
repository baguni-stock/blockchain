package keeper

import (
	"context"
	"fmt"
	"strconv"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateStockTransaction(goCtx context.Context, msg *types.MsgCreateStockTransaction) (*types.MsgCreateStockTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// address로 stock_transaction 얻기
	stock_transaction, isFound := k.GetStockTransaction(ctx, msg.Creator)
	if !isFound {
		var holding_stocks []*types.HoldingStock
		stock_transaction = types.StockTransaction{
			Creator:       msg.Creator,
			HoldingStocks: holding_stocks,
		}
	}

	//stock code의 가격(amount) 얻기
	stock_data, isFound := k.GetStockData(ctx, msg.Code)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not exists today sotck data")

	}

	amount := stock_data.Amount
	total_amount := amount * msg.Count

	// 주식보유 체크 & 구매한 주식 보유량과 매도금액 증가
	isFound = false
	var code_count int
	for code_count = 0; code_count < len(stock_transaction.HoldingStocks); code_count += 1 {
		if stock_transaction.HoldingStocks[code_count].Code == msg.Code {
			isFound = true
			break
		}
	}

	if isFound {
		stock_transaction.HoldingStocks[code_count].Count += msg.Count
		stock_transaction.HoldingStocks[code_count].PurchasAmount += total_amount
	} else {
		holding_stock := types.HoldingStock{
			Code:          msg.Code,
			Count:         msg.Count,
			PurchasAmount: total_amount,
		}
		stock_transaction.HoldingStocks = append(stock_transaction.HoldingStocks, &holding_stock)
	}

	//주식 구매
	k.SetStockTransaction(
		ctx,
		stock_transaction,
	)

	// 구매 대금 지불
	stake := strconv.FormatInt(int64(total_amount), 10) + "stake"
	burn_coin, err := sdk.ParseCoinsNormalized(stake)
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

	return &types.MsgCreateStockTransactionResponse{}, nil

}

func (k msgServer) DeleteStockTransaction(goCtx context.Context, msg *types.MsgDeleteStockTransaction) (*types.MsgDeleteStockTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	stock_transaction, isFound := k.GetStockTransaction(ctx, msg.Creator)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("creator %v not set", msg.Creator))
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Creator != stock_transaction.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// 주식을 보유 하고 있는지 체크
	isFound = false
	var code_count int
	for code_count = 0; code_count < len(stock_transaction.HoldingStocks); code_count += 1 {
		if stock_transaction.HoldingStocks[code_count].Code == msg.Code {
			isFound = true
			break
		}
	}
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("don't have stock code : %v", msg.Code))
	}

	//stock code의 가격(amount) 얻기
	stock_data, isFound := k.GetStockData(ctx, msg.Code)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "not exists today sotck data")

	}

	amount := stock_data.Amount
	total_amount := amount * msg.Count

	// 남은주식 만큼 보유량 및 매수금액 차감
	if stock_transaction.HoldingStocks[code_count].Count == msg.Count {
		stock_transaction.HoldingStocks = append(stock_transaction.HoldingStocks[:code_count], stock_transaction.HoldingStocks[code_count+1:]...)
	} else if stock_transaction.HoldingStocks[code_count].Count > msg.Count {
		stock_transaction.HoldingStocks[code_count].Count -= msg.Count
		stock_transaction.HoldingStocks[code_count].PurchasAmount -= total_amount
	} else {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("count of stock held is small. : %v", msg.Count))
	}

	//주식 판매
	k.SetStockTransaction(
		ctx,
		stock_transaction,
	)

	// 판매대금 얻기
	stake := strconv.FormatInt(int64(total_amount), 10) + "stake"
	mint_coin, err := sdk.ParseCoinsNormalized(stake)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Coin normalized faild")
	}

	user_address, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintln(err))
	}

	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(mint_coin...))
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, user_address, mint_coin)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintln(err))
	}

	return &types.MsgDeleteStockTransactionResponse{}, nil
}
