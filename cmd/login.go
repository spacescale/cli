package cmd

import (
	"github.com/spacescale/cli/internal/auth"
	"github.com/spacescale/cli/internal/output"
	"github.com/spf13/cobra"
)

func newLoginCmd(printer output.Printer) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Authenticate the CLI",
		Long:  "Authenticate the CLI with SpaceScale. This stub currently checks for SPACESCALE_TOKEN while interactive login is still under construction.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return printer.PrintLoginStub(cmd.OutOrStdout(), auth.LoadFromEnv())
		},
	}
}
