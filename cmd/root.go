package cmd

import (
	"github.com/spacescale/cli/internal/config"
	"github.com/spacescale/cli/internal/output"
	"github.com/spacescale/cli/internal/version"
	"github.com/spf13/cobra"
)

func Execute() error {
	return newRootCmd().Execute()
}

func newRootCmd() *cobra.Command {
	cfg := config.Load()
	printer := output.NewPrinter()

	rootCmd := &cobra.Command{
		Use:           "spacescale",
		Short:         "SpaceScale CLI",
		Long:          "SpaceScale CLI provides commands for working with the SpaceScale platform.",
		Version:       version.Version,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return printer.PrintRootStub(cmd.OutOrStdout(), cfg)
		},
	}

	return rootCmd
}
