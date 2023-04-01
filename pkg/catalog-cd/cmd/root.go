package cmd

import (
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/runner"

	"github.com/spf13/cobra"
	clioptions "k8s.io/cli-runtime/pkg/genericclioptions"
)

func NewRootCmd(ioStreams *clioptions.IOStreams) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "catalog-cd",
		Long: `TODO`,
	}

	cfg := config.NewConfig()

	rootCmd.AddCommand(runner.NewRunner(NewLintCmd(), cfg, ioStreams).Cmd())

	return rootCmd
}
