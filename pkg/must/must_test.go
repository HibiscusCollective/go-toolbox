package must_test

import (
	"errors"
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
		"should panic given a non-nil error": func(g gomega.Gomega) {
			g.Expect(func() {
				must.OrPanic(valOrErr(nil, errors.New("boom")))
			}).To(gomega.Panic())
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
