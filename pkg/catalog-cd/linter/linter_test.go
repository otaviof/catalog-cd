package linter

import (
	"testing"

	o "github.com/onsi/gomega"
	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"
)

func TestNewLinter(t *testing.T) {
	g := o.NewWithT(t)

	cfg := config.NewConfig()

	l, err := NewLinter(cfg, "../../../test/resources/task.yaml")
	g.Expect(err).To(o.Succeed())
	g.Expect(l).NotTo(o.BeNil())

	err = l.Enforce()
	g.Expect(err).To(o.Succeed())
}
