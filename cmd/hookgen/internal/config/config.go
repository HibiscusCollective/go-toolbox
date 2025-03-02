// Package config implements configuration parsing for the generator
package config

import (
	"errors"
	"slices"
)

// Config contains the configuration for the generator
type Config interface {
	Projects() []Project
}

type config []Project

// Projects implements Config.
func (c config) Projects() []Project {
	return []Project(c)
}

// Create creates a new config
func Create(project Project, moreProjects ...Project) (Config, error) {
	projects := make([]Project, 1, len(moreProjects)+1)
	projects[0] = project
	projects = append(projects, moreProjects...)

	projects = slices.DeleteFunc(projects, isNilOrZero)

	if len(projects) == 0 {
		return nil, FieldErrors{
			"Projects": errors.New("projects field must not be empty"),
		}.IntoError()
	}

	return config(projects), nil
}
