package cli

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/chainstock-project/blockchain/x/blockchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
)

func CmdCreateBoardComment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-board-comment [board_comment_id] [body]",
		Short: "Create a new board-comment",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			board_comment_id := cast.ToUint64(args[0])
			body := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateBoardComment(clientCtx.GetFromAddress().String(), board_comment_id, body)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateBoardComment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-board-comment [board_comment_id] [comment_id] [body]",
		Short: "Update a board-comment",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			board_comment_id := cast.ToUint64(args[0])
			comment_id := cast.ToInt64(args[1])
			body := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateBoardComment(clientCtx.GetFromAddress().String(), board_comment_id, comment_id, body)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteBoardComment() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-board-comment [board_comment_id] [comment_id]",
		Short: "Delete a board-comment",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			board_comment_id := cast.ToUint64(args[0])
			comment_id := cast.ToInt64(args[1])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteBoardComment(clientCtx.GetFromAddress().String(), board_comment_id, comment_id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
