package main

import (
	"github.com/spf13/cobra"
)

func buildRootCmd() *cobra.Command {

	// root command flags
	var opts struct {
		Version bool
	}

	cmd := &cobra.Command{
		Use:          "buffnet",
		Short:        "The Buffanetes CLI",
		SilenceUsage: true,
	}

	cmd.Flags().BoolVarP(&opts.Version, "version", "v", false, "Show the application version")

	// cmd.AddCommand(newGetCmd(cxt))
	// cmd.AddCommand(newDescribeCmd(cxt))
	// cmd.AddCommand(instance.NewProvisionCmd(cxt))
	// cmd.AddCommand(instance.NewDeprovisionCmd(cxt))
	// cmd.AddCommand(binding.NewBindCmd(cxt))
	// cmd.AddCommand(binding.NewUnbindCmd(cxt))
	// cmd.AddCommand(newSyncCmd(cxt))

	return cmd
}
