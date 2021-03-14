package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stunndard/registry/x/registry/types"
)

func CmdBuyName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy-name [name] [bid] [onsale]",
		Short: "Buys a new name",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0])
			argsBid, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return err
			}
			argsOnsale, _ := strconv.ParseBool(args[2])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBuyName(clientCtx.GetFromAddress().String(), string(argsName), argsBid, bool(argsOnsale))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-name [name] [owner] [price] [onsale]",
		Short: "Update a name",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := string(args[0])
			argsOwner := string(args[1])
			argsPrice, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return err
			}

			argsOnsale, _ := strconv.ParseBool(args[3])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateName(clientCtx.GetFromAddress().String(), string(name), string(argsOwner), argsPrice, bool(argsOnsale))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-name [name]",
		Short: "Delete a name",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			name := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteName(clientCtx.GetFromAddress().String(), name)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
