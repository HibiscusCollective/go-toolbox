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

// NewConfig creates a new config
func NewConfig(projects ...Project) (Config, error) {
	return nil, nil
}
