package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/comdex-official/comdex/x/loan/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	cmd.AddCommand(
		txRequest(),
		txApprove(),
		txRepay(),
		txCancel(),
	)

	return cmd
}

func txRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request [app-id] [pair-id] [amountIn] [amountOut] [fee] [duration]",
		Short: "create a loan request",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			appID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			pairID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			amountIn, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return err
			}

			amountOut, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return err
			}

			fee, ok := sdk.NewIntFromString(args[4])
			if !ok {
				return err
			}

			duration, ok := sdk.NewIntFromString(args[5])
			if !ok {
				return err
			}

			msg := types.NewMsgCreate(ctx.GetFromAddress().String(), appID, pairID, amountIn, amountOut, fee, duration.Int64())

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txApprove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve [loan-id]",
		Short: "approve a loan request",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			loanID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgApprove(ctx.GetFromAddress().String(), loanID)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txRepay() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "repay [loan-id]",
		Short: "repay a loan ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			loanID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRepay(ctx.GetFromAddress().String(), loanID)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func txCancel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel [loan-id]",
		Short: "Cancel a loan ",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			loanID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgCancel(ctx.GetFromAddress().String(), loanID)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
