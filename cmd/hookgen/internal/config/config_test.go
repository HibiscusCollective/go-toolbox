package config_test

import (
	"errors"
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
	"github.com/HibiscusCollective/go-toolbox/pkg/must"
)

func TestConfig(t *testing.T) {
	t.Parallel()

	scns := map[string]func(g gomega.Gomega){
		"should return an error if the fields are empty": func(g gomega.Gomega) {
			cfg, err := config.CreateConfig(nil)

			g.Expect(cfg).To(gomega.BeZero())
			g.Expect(err).To(gomega.MatchError(config.FieldErrors{
				"Projects": errors.New("projects field must not be empty"),
			}.IntoError()))
		},
		"should return a valid config": func(g gomega.Gomega) {
			cfg, err := config.CreateConfig(
				must.OrPanic(config.NewProject("test", "test", "template")),
				must.OrPanic(config.NewProject("test2", "test2", "template2")),
			)

			g.Expect(err).To(gomega.BeNil())
			g.Expect(cfg.Projects()).To(gomega.Equal([]config.Project{
				must.OrPanic(config.NewProject("test", "test", "template")),
				must.OrPanic(config.NewProject("test2", "test2", "template2")),
			}))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}
