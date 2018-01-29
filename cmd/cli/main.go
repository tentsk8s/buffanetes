package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := buildRootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}
