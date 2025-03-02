package config

import (
	"errors"
	"slices"
)

// Project contains the configuration to generate hook config files for a project
type Project struct {
	Name      string
	Path      string
	Templates []string
}

// NewProject creates a new project configuration
func NewProject(name, path string, template string, moreTemplates ...string) (Project, error) {
	templates := make([]string, 1, len(moreTemplates)+1)
	templates[0] = template

	p := Project{
		Name: name,
		Path: path,
		Templates: slices.DeleteFunc(
			append(templates, moreTemplates...),
			isEmpty,
		),
	}

	if err := p.validate(); err != nil {
		return Project{}, err
	}

	return p, nil
}

func (p *Project) validate() error {
	errs := FieldErrors{}

	if p.Name == "" {
		errs["Name"] = errors.New("name field must not be empty")
	}

	if p.Path == "" {
		errs["Path"] = errors.New("path field must not be empty")
	}

	if len(p.Templates) == 0 || slices.ContainsFunc(p.Templates, func(t string) bool { return t == "" }) {
		errs["Templates"] = errors.New("templates field must not be empty")
	}

	return errs.IntoError()
}

func isEmpty(t string) bool { return t == "" }
