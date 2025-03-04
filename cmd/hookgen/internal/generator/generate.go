// Package generator deals with orchestrating the template generation logic. It's lightweight as most logic is delegated to the template.
package generator

import (
	"errors"
	"fmt"
	"io"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
)

// FileSystemReader is an interface that defines the logic for listing files and reading a file
type FileSystemReader interface {
	ListFiles(path string) ([]string, error)
	ReadFile(path string) (io.ReadCloser, error)
}

// FileSystemWriter is an interface that defines the logic for writing a file
type FileSystemWriter interface {
	WriteFile(path string) (io.WriteCloser, error)
}

// TemplateEngine is an interface that defines the logic for applying some data to a template to generate a file
type TemplateEngine interface {
	Apply(w io.Writer, template string, data config.Project) error
}

// TemplateGenerator is a struct that encapsulates the logic for generating hook config files using a template engine
type TemplateGenerator struct {
	reader FileSystemReader
	writer FileSystemWriter
	engine TemplateEngine
}

// Create creates a new TemplateGenerator
func Create(reader FileSystemReader, writer FileSystemWriter, engine TemplateEngine) (TemplateGenerator, error) {
	// TODO: Validate not nil

	return TemplateGenerator{
		reader: reader,
		writer: writer,
		engine: engine,
	}, nil
}

// Generate generates the hook config files for the given projects
func (g TemplateGenerator) Generate(config config.Config) error {
	const errMsg = "failed to generate hook configurations"

	if config == nil {
		return fmt.Errorf("%s: %w", errMsg, MissingParametersError("config"))
	}

	var errs error
	for _, project := range config.Projects() {
		err := g.generateProjectHookFiles(project)

		errs = errors.Join(errs, err)
	}

	if errs != nil {
		return fmt.Errorf("%s: %w", errMsg, errs)
	}

	return nil
}

func (g TemplateGenerator) generateProjectHookFiles(project config.Project) error {
	var errs error
	for _, template := range project.Templates() {
		err := g.engine.Apply(io.Discard, template, project)

		errs = errors.Join(errs, TemplateExecutionError(err, template, project))
	}

	return errs
}
