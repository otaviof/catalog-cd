package cmd

import (
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/runner"

	"github.com/spf13/cobra"
	clioptions "k8s.io/cli-runtime/pkg/genericclioptions"
)

type LintCmd struct {
	cmd *cobra.Command
}

var _ runner.SubCommand = &LintCmd{}

// Cmd shares the Cobra command instance.
func (l *LintCmd) Cmd() *cobra.Command {
	return l.cmd
}

// Complete implements runner.SubCommand
func (*LintCmd) Complete(cfg *config.Config, ioStreams *clioptions.IOStreams, args []string) error {
	panic("unimplemented")
}

// Run implements runner.SubCommand
func (*LintCmd) Run(cfg *config.Config, ioStreams *clioptions.IOStreams) error {
	panic("unimplemented")
}

// Validate implements runner.SubCommand
func (*LintCmd) Validate() error {
	panic("unimplemented")
}

func NewLintCmd() runner.SubCommand {
	cmd := &cobra.Command{
		Use: "lint",
	}
	return &LintCmd{cmd: cmd}
}
