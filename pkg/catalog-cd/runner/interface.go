package runner

import (
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"

	"github.com/spf13/cobra"
)

type SubCommand interface {
	Cmd() *cobra.Command

	Complete(_ *config.Config, _ []string) error

	Validate() error

	Run(_ *config.Config) error
}
