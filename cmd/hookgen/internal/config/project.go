package config

import "errors"

// Project contains the configuration to generate hook config files for a project
type Project struct {
	Name      string
	Path      string
	Templates []string
}

// NewProject creates a new project configuration
func NewProject(_, _ string, _ string, _ ...string) (Project, error) {
	p := Project{
		Name:      "",
		Path:      "",
		Templates: nil,
	}

	return p, p.validate()
}

func (p *Project) validate() error {
	errs := FieldErrors{}

	if p.Name == "" {
		errs["Name"] = errors.New("name field must not be empty")
	}

	if p.Path == "" {
		errs["Path"] = errors.New("path field must not be empty")
	}

	if len(p.Templates) == 0 {
		errs["Templates"] = errors.New("templates field must not be empty")
	}

	return errs.IntoError()
}
