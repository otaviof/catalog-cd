package cmd

import (
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/runner"

	"github.com/spf13/cobra"
)

type LintCmd struct {
	cmd *cobra.Command
}

var _ runner.SubCommand = &LintCmd{}

// Cmd shares the Cobra command instance.
func (l *LintCmd) Cmd() *cobra.Command {
	return l.cmd
}

func (*LintCmd) Complete(_ *config.Config, _ []string) error {
	panic("unimplemented")
}

func (*LintCmd) Run(_ *config.Config) error {
	panic("unimplemented")
}

func (*LintCmd) Validate() error {
	panic("unimplemented")
}

func NewLintCmd() runner.SubCommand {
	cmd := &cobra.Command{Use: "lint"}
	return &LintCmd{cmd: cmd}
}
