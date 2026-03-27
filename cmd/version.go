package cmd

import (
	"github.com/spacescale/cli/internal/output"
	"github.com/spf13/cobra"
)

func newVersionCmd(printer output.Printer) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show CLI version information",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return printer.PrintVersion(cmd.OutOrStdout())
		},
	}
}
