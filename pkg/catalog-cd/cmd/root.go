package cmd

import (
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/runner"

	"github.com/spf13/cobra"
	tkncli "github.com/tektoncd/cli/pkg/cli"
)

func NewRootCmd(stream *tkncli.Stream) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:  "catalog-cd",
		Long: `TODO`,
	}

	cfg := config.NewConfig(stream, rootCmd.PersistentFlags())

	rootCmd.AddCommand(runner.NewRunner(cfg, NewLintCmd()).Cmd())
	rootCmd.AddCommand(runner.NewRunner(cfg, NewProbeCmd()).Cmd())

	return rootCmd
}
