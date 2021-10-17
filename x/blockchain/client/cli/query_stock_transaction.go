package cli

import (
	"context"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListStockTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-stock-transaction",
		Short: "list all stock-transaction",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllStockTransactionRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.StockTransactionAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowStockTransaction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stock-transaction [address]",
		Short: "shows a stock-transaction",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetStockTransactionRequest{
				Creator: args[0],
			}

			res, err := queryClient.StockTransaction(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
