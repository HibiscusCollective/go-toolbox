package gen_test

import (
	_ "embed"
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/HibiscusCollective/go-toolbox/scripts/hookgen/internal/gen"
	"github.com/onsi/gomega"
)

//go:embed testdata/want.yaml
var want string

type brokenExecuter struct{}

func TestGenerateFromTemplate(t *testing.T) {
	t.Parallel()

	cases := map[string]func(g gomega.Gomega){
		"should fail to generate a lefthook config file from an invalid project config json": func(g gomega.Gomega) {
			var got strings.Builder

			err := gen.New().ProjectHooks(&got, strings.NewReader("{!}"))

			g.Expect(err).To(gomega.MatchError(gomega.MatchRegexp("^" + gen.ErrParseProjectConfigMsg)))
		},
		"should fail to generate a lefthook config file when the template returns an error": func(g gomega.Gomega) {
			var got strings.Builder

			err := gen.New(gen.WithTemplate(brokenExecuter{})).ProjectHooks(&got, strings.NewReader("{}"))

			g.Expect(err).To(gomega.MatchError(gomega.MatchRegexp("^" + gen.ErrExecuteTemplateMsg)))
		},
		"should fail to generate a lefthook config file when the ProjectConfig object is incomplete": func(g gomega.Gomega) {
			var got strings.Builder

			err := gen.New().ProjectHooks(&got, strings.NewReader(`{"name": "Test Project"}`))

			g.Expect(err).To(gomega.MatchError(gomega.MatchRegexp("^" + gen.ErrParseProjectConfigMsg)))
		},
		"should generate a valid lefthook config file from template": func(g gomega.Gomega) {
			var got strings.Builder

			err := gen.New().ProjectHooks(&got, strings.NewReader(`{"name": "Test Project", "path": "test"}`))

			g.Expect(err).To(gomega.BeNil())
			g.Expect(got.String()).To(gomega.MatchYAML(want))
		},
	}

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}

func (b brokenExecuter) Execute(_ io.Writer, _ any) error {
	return newTestError()
}

func newTestError() error {
	const TEST_ERROR_MSG = "boom"

	return errors.New(TEST_ERROR_MSG)
}
