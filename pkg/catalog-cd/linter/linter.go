package linter

import (
	"errors"
	"fmt"
	"strings"

	"github.com/otaviof/catalog-cd/pkg/catalog-cd/config"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Linter enforce linting rules on the informed Tekton resource file.
type Linter struct {
	cfg *config.Config             // global configuration
	u   *unstructured.Unstructured // object instance
}

var (
	ErrInvalidWorkspace = errors.New("invalid workspace definition")
	ErrInvalidParam     = errors.New("invalid param definition")
	ErrInvalidResult    = errors.New("invalid results definition")
)

// getNestedSlice retrieve the informed path (based on fields) as a slice.
func (l *Linter) getNestedSlice(fields ...string) ([]interface{}, error) {
	slice, found, err := unstructured.NestedSlice(l.u.Object, fields...)
	if err != nil {
		return nil, err
	}
	if !found {
		l.cfg.Errorf("# WARNING: %q is not found on the resource!\n", strings.Join(fields, "."))
		return []interface{}{}, nil
	}
	return slice, nil
}

// workspaces extract and lint the workspaces.
func (l *Linter) workspaces() error {
	l.cfg.Infof("# Inspecting workspaces...\n")
	wbs, err := l.getNestedSlice("spec", "workspaces")
	if err != nil {
		return err
	}
	if err = lowercaseSliceMapLinter(wbs); err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidWorkspace, err)
	}
	l.cfg.Infof("# Workspaces are following best practices!\n")
	return nil
}

// params extract and lint the params.
func (l *Linter) params() error {
	l.cfg.Infof("# Inspecting params...\n")
	ps, err := l.getNestedSlice("spec", "params")
	if err != nil {
		return err
	}
	if err = uppercaseSliceMapLinter(ps); err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidParam, err)
	}
	l.cfg.Infof("# Params are following best practices!\n")
	return nil
}

// results extract and lint the results.
func (l *Linter) results() error {
	l.cfg.Infof("# Inspecting results...\n")
	rs, err := l.getNestedSlice("spec", "results")
	if err != nil {
		return err
	}
	if err = uppercaseSliceMapLinter(rs); err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidResult, err)
	}
	l.cfg.Infof("# Results are following best practices!\n")
	return nil
}

// Enforce lint workspaces, params and results.
func (l *Linter) Enforce() error {
	err := l.workspaces()
	if err != nil {
		return err
	}
	if err = l.params(); err != nil {
		return err
	}
	return l.results()
}

// NewLinter instantiate the resource linter by reading and decoding the resource file.
func NewLinter(cfg *config.Config, resource string) (*Linter, error) {
	cfg.Infof("# Linting resource file %q...\n", resource)
	u, err := readAndDecodeResourceFile(resource)
	if err != nil {
		return nil, err
	}
	cfg.Infof("# Name=%q, APIVersion=%q, Kind=%q\n", u.GetName(), u.GetAPIVersion(), u.GetKind())

	return &Linter{
		cfg: cfg,
		u:   u,
	}, nil
}
