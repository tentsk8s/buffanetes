package main

import (
	"io/ioutil"
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
			metaBytes, err := ioutil.ReadFile(".buffanetes/meta.toml")
			if err != nil {
				return err
			}
			meta, err := config.ParseMeta(metaBytes)
			if err != nil {
				return nil
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
