package main

import (
	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "buffnet",
		Short: "The Buffanetes CLI",
	}

	cmd.AddCommand(envCmd())
	cmd.AddCommand(deployRootCmd())
	cmd.AddCommand(initCmd())

	return cmd
}
