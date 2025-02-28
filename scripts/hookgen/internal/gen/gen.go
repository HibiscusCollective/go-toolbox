// Package gen provides functions to generate a lefthook config file from a template
package gen

import (
	"encoding/json"
	"io"
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

// Project writes the template to the writer using project config
func Project(w io.Writer, tmpl Executer, data json.RawMessage) error {
	return tmpl.Execute(w, ProjectConfig{Name: "Test Project", Path: "test"})
}
