package generator_test

import (
	_ "embed"
	"html/template"
	"io"
	"os"
	"path"

	"errors"
	"fmt"
	"testing"

	"github.com/onsi/gomega"
	"github.com/spf13/afero"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/filesys"
	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/generator"
	"github.com/HibiscusCollective/go-toolbox/pkg/fxslice"
	"github.com/HibiscusCollective/go-toolbox/pkg/must"
)

func TestGenerator(t *testing.T) {
	t.Parallel()

	const errMsg = "failed to generate hook configurations"

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should return an error if the config is nil": func(t testing.TB, g gomega.Gomega) {
			gen, err := generator.Create(filesys.New(afero.NewMemMapFs()), fakeEngine{})
			g.Expect(err).To(gomega.BeNil())

			err = gen.Generate(nil)

			g.Expect(err).To(gomega.MatchError(fmt.Errorf("%s: %w", errMsg, generator.MissingParametersError("config"))))
		},
		"should return an error if the template execution fails": func(t testing.TB, g gomega.Gomega) {
			gen, err := generator.Create(filesys.New(afero.NewMemMapFs()), newBrokenEngine(errors.New("boom")))
			g.Expect(err).To(gomega.BeNil())

			cfg := projects{
				{"test project 1", "test/project1", []string{"test.tmpl"}},
			}.mustCreateConfig(t)

			err = gen.Generate(cfg)

			g.Expect(err).To(gomega.MatchError(fmt.Errorf(
				"%s: %w",
				errMsg,
				errors.Join(nil, generator.TemplateExecutionError(errors.New("boom"), "test.tmpl", cfg.Projects()[0])),
			)))
		},
		"should write a file for each template configured for the project": func(t testing.TB, g gomega.Gomega) {
			tfs := afero.NewMemMapFs()

			cfg := projects{
				{"test project 1", "test/project1", []string{"test1.txt.tmpl", "test2.txt.tmpl"}},
			}.mustCreateConfig(t)

			gen := generator.MustCreate(t, filesys.New(tfs), fakeEngine{
				"test1.txt.tmpl": "hello, {{.Name}}!",
				"test2.txt.tmpl": "goodbye, {{.Name}}!",
			})

			err := gen.Generate(cfg)
			g.Expect(err).To(gomega.BeNil())

			g.Expect(mustReadFilesInDir(t, tfs, "test/project1")).To(gomega.Equal(testFiles{
				"test1.txt": "hello, test project 1!",
				"test2.txt": "goodbye, test project 1!",
			}))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(t, gomega.NewWithT(t))
		})
	}
}

type project struct {
	name      string
	path      string
	templates []string
}

type projects []project

func (ps projects) mustCreateConfig(t testing.TB) config.Config {
	t.Helper()

	if len(ps) < 1 {
		t.Fatal("at least one project must be provided")
	}

	projects, err := fxslice.TryTransform(ps, func(p project) (config.Project, error) {
		return config.CreateProject(p.name, p.path, p.templates[0], p.templates[1:]...)
	})
	if err != nil {
		t.Fatal("unexpected error(s) creating projects: %w", err)
	}

	return must.Succeed(config.Create(projects[0], projects[1:]...)).OrFailTest(t)
}

type testFiles map[string]string

func mustReadFilesInDir(t testing.TB, fs afero.Fs, dir string) testFiles {
	t.Helper()

	dirFileNames := fxslice.Transform(
		must.Succeed(afero.ReadDir(fs, dir)).OrFailTest(t),
		func(fi os.FileInfo) string { return fi.Name() },
	)

	content := must.Succeed(fxslice.TryTransform(dirFileNames, func(fn string) ([]byte, error) {
		f, err := fs.Open(path.Join(dir, fn))
		if err != nil {
			return nil, err
		}

		return io.ReadAll(f)
	})).OrFailTest(t)

	tfs := make(testFiles, len(dirFileNames))
	for i, file := range dirFileNames {
		tfs[file] = string(content[i])
	}

	return tfs
}

type fakeEngine map[string]string

type brokenEngine struct {
	err error
}

func newBrokenEngine(err error) brokenEngine {
	return brokenEngine{
		err: err,
	}
}

func (f fakeEngine) Apply(w io.Writer, tmplFile string, data config.Project) error {
	tmpl, err := template.New(tmplFile).Parse(f[tmplFile])
	if err != nil {
		return err
	}

	return tmpl.Execute(w, data)
}

func (f brokenEngine) Apply(_ io.Writer, _ string, _ config.Project) error {
	return f.err
}
