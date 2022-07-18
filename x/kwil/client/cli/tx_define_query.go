package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"kwil/x/kwil/types"
)

var _ = strconv.Itoa(0)

func CmdDefineQuery() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "define-query [par-quer] [publicity]",
		Short: "Broadcast message DefineQuery",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argParQuer := args[0]
			argPublicity, err := cast.ToBoolE(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDefineQuery(
				clientCtx.GetFromAddress().String(),
				argParQuer,
				argPublicity,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
