package config

import (
	"errors"
	"slices"
)

// Project contains the configuration to generate hook config files for a project
type Project interface {
	Name() string
	Path() string
	Templates() []string
}

type project struct {
	name      string
	path      string
	templates []string
}

// NewProject creates a new project configuration
func NewProject(name, path string, template string, moreTemplates ...string) (Project, error) {
	templates := make([]string, 1, len(moreTemplates)+1)
	templates[0] = template

	p := project{
		name: name,
		path: path,
		templates: slices.DeleteFunc(
			append(templates, moreTemplates...),
			isEmpty,
		),
	}

	if err := p.validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p project) Name() string {
	return p.name
}

func (p project) Path() string {
	return p.path
}

func (p project) Templates() []string {
	return p.templates
}

func (p project) validate() error {
	errs := FieldErrors{}

	if p.name == "" {
		errs["Name"] = errors.New("name field must not be empty")
	}

	if p.path == "" {
		errs["Path"] = errors.New("path field must not be empty")
	}

	if len(p.templates) == 0 {
		errs["Templates"] = errors.New("templates field must not be empty")
	}

	return errs.IntoError()
}

func isEmpty(t string) bool { return t == "" }
