package main

import (
	"errors"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/tentsk8s/buffanetes/pkg/config"
	"github.com/tentsk8s/buffanetes/pkg/web"
)

func deployGriftCmd(cl *helm.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "grift [name]",
		Short:   "Run a grift in Kubernetes",
		Example: "buffnet deploy grift mygrift",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("whoops, you forgot to add a grift name")
			}
			chart := grift.BuildHelmChart(cfg, orgAndApp)
			color.Green("built helm chart")
			if err := cl.InstallRelease(chart); err != nil {
				return err
			}
			color.Green("installed helm release")
			// wait for the job to complete
			color.Green("grift complete!")
			// get logs from job
			// delete job
			return nil
		},
	}
	return cmd
}
