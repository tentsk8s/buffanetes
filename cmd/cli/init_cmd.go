package main

import (
	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize your Buffalo app to work with Buffanetes",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	return cmd
}
