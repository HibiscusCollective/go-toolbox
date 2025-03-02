package must_test

import (
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/pkg/must"
)

func TestMustOrPanic(t *testing.T) {
	t.Parallel()

	scns := map[string]func(g gomega.Gomega){
		"should not panic given a nil error": func(g gomega.Gomega) {
			val := must.OrPanic(valOrErr("hello, world!", nil))

			g.Expect(val).To(gomega.Equal("hello, world!"))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}

func valOrErr(val any, err error) (any, error) {
	return val, err
}
