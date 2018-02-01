package main

import (
	"github.com/spf13/cobra"
)

func buildRootCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "buffnet",
		Short: "The Buffanetes CLI",
	}

	cmd.AddCommand(envCmd())

	return cmd
}
