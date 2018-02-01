package main

import (
	"github.com/spf13/cobra"
)

func deployRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy something in your Buffalo app to Kubernetes",
	}
	return cmd
}
