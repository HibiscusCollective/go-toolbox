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
	Apply(template string, data config.Project) (string, error)
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
	err := validateArgs(config)
	if err != nil {
		return fmt.Errorf("failed to generate hook configurations: %w", err)
	}

	return nil
}

func validateArgs(config config.Config) error {
	var errs []error

	if config == nil || len(config.Projects()) == 0 {
		errs = append(errs, fmt.Errorf("config argument is required"))
	}

	return errors.Join(errs...)
}
