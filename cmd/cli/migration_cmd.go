package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tentsk8s/pkg/migration"
)

func migrationCmd(cl *helm.Client) *cobra.Command {
	return &cobra.Command{
		Use:     "migration [up/down]",
		Short:   "Deploy your Buffanetes web server to Kubernetes",
		Example: "buffnet deploy web myorg/myapp",
		RunE: func(cmd *cobra.Command, args []string) error {
			direction := "up"
			if len(args) > 0 {
				direction = args[1]
			}
			if direction != "up" && direction != "down" {
				return fmt.Errorf("the only migrations are 'up' or 'down'")
			}
			if direction == "down" {
				// TODO: ask "are you sure???"
			}
			// TODO: build image
			chart := migration.BuildChart()
			if err := cl.InstallRelease(chart); err != nil {
				return err
			}
			color.Green("migration started!")
			// TODO: watch job
			// TODO: get job logs & print them out
			// TODO: delete job? or do something so that helm upgrade can work on the job later?
			return nil

		},
	}
}
