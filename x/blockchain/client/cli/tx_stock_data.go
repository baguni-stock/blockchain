package cli

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdCreateStockData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-stock-data [date] [stock_type] [stocks]",
		Short: "Create a new stock-data",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			date := args[0]
			stock_type := args[1]
			var stocks []*types.Stock
			for i := 2; i < len(args); i++ {
				stock_split := strings.Split(args[i], "=")
				stock := types.Stock{
					Code:   stock_split[0],
					Amount: stock_split[1],
				}
				stock_pointer := &stock
				stocks = append(stocks, stock_pointer)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateStockData(clientCtx.GetFromAddress().String(), date, stock_type, stocks)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateStockData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-stock-data [date] [stocks]",
		Short: "Update a stock-data",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			date := args[0]
			stock_type := args[1]

			var stocks []*types.Stock
			for i := 1; i < len(args); i++ {
				stock_split := strings.Split(args[i], "=")
				stock := types.Stock{
					Code:   stock_split[0],
					Amount: stock_split[1],
				}
				stock_pointer := &stock
				stocks = append(stocks, stock_pointer)
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateStockData(clientCtx.GetFromAddress().String(), date, stock_type, stocks)
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
		Use:   "delete-stock-data [date]",
		Short: "Delete a stock-data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			date := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteStockData(clientCtx.GetFromAddress().String(), date)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
