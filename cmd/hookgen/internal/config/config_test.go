package config_test

import (
	"errors"
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	scns := map[string]func(g gomega.Gomega){
		"should return an error if the fields are empty": func(g gomega.Gomega) {
			cfg, err := config.CreateConfig()

			g.Expect(cfg).To(gomega.BeZero())
			g.Expect(err).To(gomega.MatchError(config.FieldErrors{
				"Projects": errors.New("projects field must not be empty"),
			}.IntoError()))
		},
		"should return a valid config": func(g gomega.Gomega) {
			cfg, err := config.CreateConfig(
				config.NewProject("test", "test", "template"),
				config.NewProject("test2", "test2", "template2"),
			)
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}
