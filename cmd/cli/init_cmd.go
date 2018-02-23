package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func initCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize your Buffalo app to work with Buffanetes",
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(".buffanetes"); err == nil {
				return fmt.Errorf("the .buffanetes directory already exists, bailing out")
			}
			curDir, err := os.Getwd()
			if err != nil {
				return err
			}
			appName := filepath.Base(curDir)
			logger.Printf("initializing buffanetes app %s", appName)
			return nil
		},
	}
	return cmd
}
