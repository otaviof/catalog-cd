package runner

import (
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"

	"github.com/spf13/cobra"
)

type Runner struct {
	cfg        *config.Config
	subCommand SubCommand
}

func (r *Runner) Cmd() *cobra.Command {
	return r.subCommand.Cmd()
}

func (r *Runner) RunE(_ *cobra.Command, args []string) error {
	err := r.subCommand.Complete(r.cfg, args)
	if err != nil {
		return err
	}
	if err = r.subCommand.Validate(); err != nil {
		return err
	}
	return r.subCommand.Run(r.cfg)
}

// NewRunner instantiates a Runner, making sure it's RunE command is mapped to the local method,
// which executes the whole interface workflow.
func NewRunner(cfg *config.Config, subCommand SubCommand) *Runner {
	r := &Runner{
		cfg:        cfg,
		subCommand: subCommand,
	}
	r.subCommand.Cmd().RunE = r.RunE
	return r
}
