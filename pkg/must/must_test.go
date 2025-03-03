package must_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/pkg/must"
)

func TestMustOrPanic(t *testing.T) {
	t.Parallel()

	scns := map[string]func(g gomega.Gomega){
		"should not panic given a nil error": func(g gomega.Gomega) {
			val := must.Succeed(valOrErr("hello, world!", nil)).OrPanic()

			g.Expect(val).To(gomega.Equal("hello, world!"))
		},
		"should panic given a non-nil error": func(g gomega.Gomega) {
			g.Expect(func() {
				must.Succeed(valOrErr(nil, errors.New("boom"))).OrPanic()
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

func TestOrFail(t *testing.T) {
	t.Parallel()

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should not fail given a nil error": func(t testing.TB, g gomega.Gomega) {
			val := must.Succeed(valOrErr("hello, world!", nil)).OrFail(t)

			g.Expect(val).To(gomega.Equal("hello, world!"))
		},
		"should fail given a non-nil error": func(t testing.TB, g gomega.Gomega) {
			tester := mockT(t)
			must.Succeed(valOrErr(nil, errors.New("boom"))).OrFail(tester)

			g.Expect(tester.fatal).To(gomega.Equal("unexpected error: boom"))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(t, gomega.NewWithT(t))
		})
	}
}

type tm struct {
	testing.TB
	fatal string
}

func valOrErr(val any, err error) (any, error) {
	return val, err
}

func mockT(t testing.TB) *tm {
	return &tm{TB: t}
}

func (t *tm) Fatalf(format string, args ...any) {
	t.fatal = fmt.Sprintf(format, args...)
}
