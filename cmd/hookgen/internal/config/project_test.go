package config_test

import (
	"errors"
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
)

func TestProject(t *testing.T) {
	t.Parallel()

	scns := map[string]func(g gomega.Gomega){
		"should return an error constructing an empty project": func(g gomega.Gomega) {
			project, err := config.NewProject("", "", "")

			g.Expect(project).To(gomega.BeZero())

			g.Expect(err).To(gomega.MatchError(config.FieldErrors{
				"Name":      errors.New("name field must not be empty"),
				"Path":      errors.New("path field must not be empty"),
				"Templates": errors.New("templates field must not be empty"),
			}.IntoError()))
		},
		"should return a valid project": func(g gomega.Gomega) {
			project, err := config.NewProject("test_project", "test", "template1", "template2")

			g.Expect(err).To(gomega.BeNil())
			g.Expect(project.Name()).To(gomega.Equal("test_project"))
			g.Expect(project.Path()).To(gomega.Equal("test"))
			g.Expect(project.Templates()).To(gomega.Equal([]string{"template1", "template2"}))
		},
		"should filter out empty strings from the templates list": func(g gomega.Gomega) {
			project, err := config.NewProject("test_project", "test", "template1", "", "template2")

			g.Expect(err).To(gomega.BeNil())
			g.Expect(project.Templates()).To(gomega.Equal([]string{"template1", "template2"}))
		},
		"should return an error if the templates list is filtered to empty": func(g gomega.Gomega) {
			project, err := config.NewProject("test_project", "test", "", "")

			g.Expect(project).To(gomega.BeZero())
			g.Expect(err).To(gomega.MatchError(config.FieldErrors{
				"Templates": errors.New("templates field must not be empty"),
			}.IntoError()))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}
