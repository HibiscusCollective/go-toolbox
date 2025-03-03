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

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should return an error if the fields are empty": func(t testing.TB, g gomega.Gomega) {
			cfg, err := config.Create(nil)

			g.Expect(cfg).To(gomega.BeZero())
			g.Expect(err).To(gomega.MatchError(config.FieldErrors{
				"Projects": errors.New("projects field must not be empty"),
			}.IntoError()))
		},
		"should return a valid config": func(t testing.TB, g gomega.Gomega) {
			cfg, err := config.Create(
				must.Succeed(config.CreateProject("test", "test", "template")).OrFail(t),
				must.Succeed(config.CreateProject("test2", "test2", "template2")).OrFail(t),
			)

			g.Expect(err).To(gomega.BeNil())
			g.Expect(cfg.Projects()).To(gomega.Equal([]config.Project{
				must.Succeed(config.CreateProject("test", "test", "template")).OrFail(t),
				must.Succeed(config.CreateProject("test2", "test2", "template2")).OrFail(t),
			}))
		},
		"should filter out empty projects from the config": func(t testing.TB, g gomega.Gomega) {
			cfg, err := config.Create(
				must.Succeed(config.CreateProject("test", "test", "template")).OrFail(t),
				config.ZeroProject(),
				must.Succeed(config.CreateProject("test2", "test2", "template2")).OrFail(t),
			)

			g.Expect(err).To(gomega.BeNil())
			g.Expect(cfg.Projects()).To(gomega.Equal([]config.Project{
				must.Succeed(config.CreateProject("test", "test", "template")).OrFail(t),
				must.Succeed(config.CreateProject("test2", "test2", "template2")).OrFail(t),
			}))
		},
		"should filter out nil projects from the config": func(t testing.TB, g gomega.Gomega) {
			cfg, err := config.Create(
				must.Succeed(config.CreateProject("test", "test", "template")).OrFail(t),
				nil,
				must.Succeed(config.CreateProject("test2", "test2", "template2")).OrFail(t),
			)

			g.Expect(err).To(gomega.BeNil())
			g.Expect(cfg.Projects()).To(gomega.Equal([]config.Project{
				must.Succeed(config.CreateProject("test", "test", "template")).OrFail(t),
				must.Succeed(config.CreateProject("test2", "test2", "template2")).OrFail(t),
			}))
		},
		"should return an error if the projects list is filtered to empty": func(t testing.TB, g gomega.Gomega) {
			cfg, err := config.Create(
				nil,
				config.ZeroProject(),
			)

			g.Expect(cfg).To(gomega.BeZero())
			g.Expect(err).To(gomega.MatchError(config.FieldErrors{
				"Projects": errors.New("projects field must not be empty"),
			}.IntoError()))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(t, gomega.NewWithT(t))
		})
	}
}
