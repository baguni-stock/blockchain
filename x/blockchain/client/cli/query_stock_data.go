package cli

import (
	"context"

	"errors"
	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdListStockData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-stock-data",
		Short: "list all stock-data",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllStockDataRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.StockDataAll(context.Background(), params)
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

func CmdShowStockData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stock-data [date]",
		Short: "shows a stock-data",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetStockDataRequest{
				Date: args[0],
			}

			res, err := queryClient.StockData(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowStockDataCode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-stock-data-code [date] [code]",
		Short: "shows a stock-data",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetStockDataRequest{
				Date: args[0],
			}

			res, err := queryClient.StockData(context.Background(), params)
			if err != nil {
				return err
			}
			stocks := res.StockData.Stocks
			for i := 0; i < len(stocks); i++ {
				if stocks[i].Code == args[1] {
					println(stocks[i].Amount)
					return nil
				}

			}
			return errors.New("can't find code")
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
