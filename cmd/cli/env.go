package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tentsk8s/buffanetes/pkg/config"
)

func envCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "env",
		Short: "Get buffanetes environment variables",
		RunE: func(cmd *cobra.Command, args []string) error {
			meta := new(config.Meta)
			if err := config.ParseFile("meta.toml", meta); err != nil {
				return err
			}
			color.Green("Environment Variables")
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Buffanetes Name", "Env Var Name"})
			for k, v := range meta.Env {
				table.Append([]string{k, v})
			}
			table.Render()
			return nil
		},
	}
}
