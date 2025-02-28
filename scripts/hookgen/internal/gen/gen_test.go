package gen_test

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/HibiscusCollective/go-toolbox/scripts/hookgen/internal/gen"
	"github.com/onsi/gomega"
)

//go:embed testdata/want.yaml
var want string

func TestGenerateFromTemplate(t *testing.T) {
	t.Parallel()

	cases := map[string]func(g gomega.Gomega){
		"should generate a valid lefthook config file from template": func(g gomega.Gomega) {
			var got strings.Builder
			tmpl := gen.Templates()

			err := gen.Project(&got, tmpl, []byte(`{"name": "Test Project", "path": "test"}`))

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
