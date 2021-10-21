package cli

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdCreateStockData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-stock-data [code] [market_type] [amount] [date]",
		Short: "Create a new stock-data",
		Args:  cobra.MinimumNArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var stocks []*types.StockData
			creator := clientCtx.GetFromAddress().String()
			for i := 0; i < len(args); i += 4 {
				code := args[i]
				market_type := args[i+1]
				amount, err := cast.ToInt32E(args[i+2])
				if err != nil {
					return err
				}
				date := args[i+3]
				stock := types.StockData{
					Creator:    creator,
					Code:       code,
					MarketType: market_type,
					Amount:     amount,
					Date:       date,
				}
				fmt.Printf("%v\n", stock)
				stocks = append(stocks, &stock)
			}

			msg := types.NewMsgCreateStockData(creator, stocks)

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteStockData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-stock-data [code]",
		Short: "Delete a stock-data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			code := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteStockData(clientCtx.GetFromAddress().String(), code)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
