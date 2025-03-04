// Package generator deals with orchestrating the template generation logic. It's lightweight as most logic is delegated to the template.
package generator

import (
	"errors"
	"fmt"
	"io"
	"path"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/filesys"
)

// FSCreator is an interface that defines the logic for creating a writeable file and for creating a directory
type FSCreator interface {
	filesys.FileCreator
	filesys.DirMaker
}

// TemplateEngine is an interface that defines the logic for applying some data to a template to generate a file
type TemplateEngine interface {
	Apply(w io.Writer, template string, data config.Project) error
}

// TemplateGenerator is a struct that encapsulates the logic for generating hook config files using a template engine
type TemplateGenerator struct {
	fsc    FSCreator
	engine TemplateEngine
}

// Create creates a new TemplateGenerator
func Create(fsc FSCreator, engine TemplateEngine) (TemplateGenerator, error) {
	// TODO: Validate not nil

	return TemplateGenerator{
		fsc:    fsc,
		engine: engine,
	}, nil
}

// Generate generates the hook config files for the given projects
func (g TemplateGenerator) Generate(config config.Config) error {
	const errMsg = "failed to generate hook configurations"

	if config == nil {
		return fmt.Errorf("%s: %w", errMsg, MissingParametersError("config"))
	}

	errs := g.generateAllProjectHooks(config.Projects()...)
	if errs != nil {
		return fmt.Errorf("%s: %w", errMsg, errs)
	}

	return nil
}

func (g TemplateGenerator) generateAllProjectHooks(project ...config.Project) error {
	var errs error

	for _, project := range project {
		if err := g.fsc.MkdirAll(path.Dir(project.Path())); err != nil {
			// errs = errors.Join(errs, err)
			panic(err) // TODO: Handle this better
		}

		if err := g.generateProjectHookFiles(project); err != nil {
			genErrs := err.(interface{ Unwrap() []error }).Unwrap()
			for _, err := range genErrs {
				errs = errors.Join(errs, err)
			}
		}
	}

	return errs
}

func (g TemplateGenerator) generateProjectHookFiles(project config.Project) error {
	var errs error

	for _, template := range project.Templates() {
		filename := path.Base(template)
		filename = string([]rune(filename)[:len(filename)-len(path.Ext(filename))])

		f, err := g.fsc.Create(path.Join(project.Path(), filename))
		if err != nil {
			// errs = errors.Join(errs, err)
			panic(err) // TODO: Handle this better
		}

		if err := g.engine.Apply(f, template, project); err != nil {
			errs = errors.Join(errs, TemplateExecutionError(err, template, project))
		}
	}

	return errs
}
