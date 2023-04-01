package runner

import (
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	clioptions "k8s.io/cli-runtime/pkg/genericclioptions"
)

type Runner struct {
	subCommand SubCommand
	cfg        *config.Config
	ioStreams  *genericclioptions.IOStreams
}

func (r *Runner) Cmd() *cobra.Command {
	return r.subCommand.Cmd()
}

func (r *Runner) RunE(_ *cobra.Command, args []string) error {
	err := r.subCommand.Complete(r.cfg, r.ioStreams, args)
	if err != nil {
		return err
	}
	if err = r.subCommand.Validate(); err != nil {
		return err
	}
	return r.subCommand.Run(r.cfg, r.ioStreams)
}

func NewRunner(subCommand SubCommand, cfg *config.Config, ioStreams *clioptions.IOStreams) *Runner {
	return &Runner{
		subCommand: subCommand,
		cfg:        cfg,
		ioStreams:  ioStreams,
	}
}
