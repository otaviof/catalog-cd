package runner

import (
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"

	"github.com/spf13/cobra"
	clioptions "k8s.io/cli-runtime/pkg/genericclioptions"
)

type SubCommand interface {
	Cmd() *cobra.Command

	Complete(_ *config.Config, _ *clioptions.IOStreams, _ []string) error

	Validate() error

	Run(_ *config.Config, _ *clioptions.IOStreams) error
}
