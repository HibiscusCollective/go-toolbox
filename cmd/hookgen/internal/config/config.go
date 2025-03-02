// Package config implements configuration parsing for the generator
package config

import "errors"

// Config contains the configuration for the generator
type Config interface {
	Projects() []Project
}

type config []Project

// Projects implements Config.
func (c config) Projects() []Project {
	return []Project(c)
}

// CreateConfig creates a new config
func CreateConfig(project Project, moreProjects ...Project) (Config, error) {
	if project == nil {
		return nil, FieldErrors{
			"Projects": errors.New("projects field must not be empty"),
		}.IntoError()
	}

	projects := make([]Project, 1, len(moreProjects)+1)
	projects[0] = project

	if len(moreProjects) > 0 {
		projects = append(projects, moreProjects...)
	}

	return config(projects), nil
}
