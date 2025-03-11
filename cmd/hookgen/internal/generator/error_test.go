package generator_test

import (
	"errors"
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
	generator "github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/generator"
	"github.com/HibiscusCollective/go-toolbox/pkg/fxslice"
	"github.com/HibiscusCollective/go-toolbox/pkg/must"
)

func TestInvalidParametersError(t *testing.T) {
	t.Parallel()

	scns := map[string]func(tb testing.TB, g gomega.Gomega){
		"should be nil if the parameters are empty": func(tb testing.TB, g gomega.Gomega) {
			err := generator.MissingParametersError("")

			g.Expect(err).To(gomega.BeNil())
		},
		"should return an error given a parameter name": func(tb testing.TB, g gomega.Gomega) {
			err := generator.MissingParametersError("Name")

			g.Expect(err).To(gomega.MatchError(errors.Join(
				generator.ParameterErrorImpl{
					Msg:   "missing required parameter: Name",
					Label: "Name",
				},
			)))
		},
		"should return an error given multiple parameters": func(tb testing.TB, g gomega.Gomega) {
			err := generator.MissingParametersError("Name", "Path")

			g.Expect(err).To(gomega.MatchError(errors.Join(
				generator.ParameterErrorImpl{
					Msg:   "missing required parameter: Name",
					Label: "Name",
				},
				generator.ParameterErrorImpl{
					Msg:   "missing required parameter: Path",
					Label: "Path",
				},
			)))
		},
		"should extract the parameter name from the error": func(tb testing.TB, g gomega.Gomega) {
			err := generator.MissingParametersError("Name")

			var unwrapper interface{ Unwrap() []error }
			if !errors.As(err, &unwrapper) {
				tb.Errorf("error is not a joined error: %v", err)
			}

			errs := must.GetOrFailTest(fxslice.Cast[error, generator.ParameterError](unwrapper.Unwrap()))(tb)
			g.Expect(errs).To(gomega.HaveLen(1))
			g.Expect(errs[0].Parameter()).To(gomega.Equal("Name"))
		},
		"should print 'missing required parameter: Name' as an error message": func(tb testing.TB, g gomega.Gomega) {
			err := generator.MissingParametersError("Name")

			g.Expect(err).To(gomega.MatchError("missing required parameter: Name"))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(t, gomega.NewWithT(t))
		})
	}
}

func TestTemplateError(t *testing.T) {
	t.Parallel()

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should be nil if the error is nil": func(t testing.TB, g gomega.Gomega) {
			err := generator.TemplateExecutionError(nil, "test.tmpl", must.GetOrFailTest(config.CreateProject("test", "test", "template"))(t))
			g.Expect(err).To(gomega.BeNil())
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(t, gomega.NewWithT(t))
		})
	}
}
