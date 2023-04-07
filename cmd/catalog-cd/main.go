package main

import (
	"fmt"
	"os"

	"github.com/otaviof/catalog-cd/pkg/catalog-cd/cmd"
	tkncli "github.com/tektoncd/cli/pkg/cli"
)

func main() {
	rootCmd := cmd.NewRootCmd(&tkncli.Stream{
		In:  os.Stdin,
		Out: os.Stdout,
		Err: os.Stderr,
	})
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
