package main

import (
	"os"
)

func main() {
	cmd := rootCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}

}
