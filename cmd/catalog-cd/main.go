package main

import (
	"fmt"
	"os"

	"github.com/otaviof/catalog-cd/pkg/catalog-cd/cmd"
	clioptions "k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	ioStreams := &clioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr}
	rootCmd := cmd.NewRootCmd(ioStreams)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
