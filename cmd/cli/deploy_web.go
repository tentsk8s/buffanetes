package main

import (
	"errors"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tentsk8s/buffanetes/pkg/config"
	"github.com/tentsk8s/buffanetes/pkg/web"
)

func deployWebCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "web",
		Short:   "Deploy your Buffanetes web server to Kubernetes",
		Example: "buffnet deploy web myorg/myapp",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("you forgot to add an org/app")
			}
			orgAndApp, err := config.ParseOrgAndApp(args[0])
			if err != nil {
				return err
			}
			cfg := new(config.Web)
			if err := config.ParseFile("web.toml", cfg); err != nil {
				return err
			}
			chart := web.BuildHelmChart(cfg, orgAndApp)
			color.Green(chart)
			return nil
		},
	}
	return cmd
}
