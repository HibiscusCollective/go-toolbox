package generator_test

import (
	"embed"
	_ "embed"
	"io"

	"errors"
	"fmt"
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/generator"
)

//go:embed testdata/*
var testFS embed.FS

type stubReader map[string]string

type stubWriter map[string]io.Writer

type stubEngine struct {
	err error
}

func TestGenerator(t *testing.T) {
	t.Parallel()

	const errMsg = "failed to generate hook configurations"

	scns := map[string]func(g gomega.Gomega){
		"should return an error if the config is nil": func(g gomega.Gomega) {
			gen, err := generator.Create(stubReader{}, stubWriter{}, stubEngine{})
			g.Expect(err).To(gomega.BeNil())
			g.Expect(gen).To(gomega.Not(gomega.BeNil()))

			err = gen.Generate(nil)

			g.Expect(err).To(gomega.MatchError(fmt.Errorf("%s: %w", errMsg, generator.ParameterErrors{
				"cwd":    errors.New("cwd argument is required"),
				"config": errors.New("config argument is required"),
			}.IntoError())))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}

func (s stubReader) ListFiles(path string) ([]string, error) {
	panic("unimplemented")
}

func (s stubReader) ReadFile(path string) (io.ReadCloser, error) {
	panic("unimplemented")
}

func (s stubWriter) WriteFile(path string) (io.WriteCloser, error) {
	panic("unimplemented")
}

func (s stubEngine) Apply(template string, data config.Project) (string, error) {
	panic("unimplemented")
}
