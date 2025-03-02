// Package gen deals with orchestrating the template generation logic. It's lightweight as most logic is delegated to the template.
package gen

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
)

// Gen generates the hook config files for the given projects
func Gen(cwd fs.FS, config config.Config) error {
	err := validateArgs(cwd, config)
	if err != nil {
		return fmt.Errorf("failed to generate hook configurations: %w", err)
	}

	return nil
}

func validateArgs(cwd fs.FS, config config.Config) error {
	errs := ParameterErrors{}
	if cwd == nil {
		errs["cwd"] = errors.New("cwd argument is required")
	}

	if config == nil || len(config.Projects()) == 0 {
		errs["config"] = errors.New("config argument is required")
	}

	return errs.IntoError()
}
