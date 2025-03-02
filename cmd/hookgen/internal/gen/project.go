package gen

import (
	"encoding/json"
	"fmt"
	"io"
)

// ProjectConfig is the configuration for a project
type ProjectConfig struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func DecodeProjectConfig(r io.Reader) (ProjectConfig, error) {
	var cfg ProjectConfig
	if err := json.NewDecoder(r).Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("%s: %w", errParseProjectConfigMsg, err)
	}
	return cfg, nil
}

func (c ProjectConfig) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("name is required")
	}
	if c.Path == "" {
		return fmt.Errorf("path is required")
	}
	return nil
}
