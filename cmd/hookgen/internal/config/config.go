package config

// Config contains the configuration for the generator
type Config interface {
	Projects() []Project
}

type config []Project

// Projects implements Config.
func (c config) Projects() []Project {
	panic("unimplemented")
}

// CreateConfig creates a new config
func CreateConfig(projects ...Project) (Config, error) {
	return nil, nil
}
