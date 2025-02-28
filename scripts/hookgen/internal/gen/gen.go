// Package gen provides functions to generate a lefthook config file from a template
package gen

import (
	"encoding/json"
	"fmt"
	"io"
)

const (
	errParseProjectConfigMsg = "failed to parse project config: "
	errExecuteTemplateMsg    = "failed to execute template: "
)

// Executer represents an object that can write a template to an io.Write
type Executer interface {
	Execute(io.Writer, any) error
}

// ProjectConfig is the configuration for a project
type ProjectConfig struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// Generator generates a lefthook config file
type Generator struct {
	tmpl Executer
}

// New returns a new Generator
func New(opts ...func(*Generator)) Generator {
	g := Generator{
		tmpl: Templates(),
	}

	for _, opt := range opts {
		opt(&g)
	}

	return g
}

// ProjectHooks writes the template to the writer using project config
func (g Generator) ProjectHooks(w io.Writer, r io.Reader) error {
	var cfg ProjectConfig

	if err := json.NewDecoder(r).Decode(&cfg); err != nil {
		return fmt.Errorf("%s: %w", errParseProjectConfigMsg, err)
	}

	if err := g.tmpl.Execute(w, cfg); err != nil {
		return fmt.Errorf("%s: %w", errExecuteTemplateMsg, err)
	}

	return nil
}
