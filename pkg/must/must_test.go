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
			val := must.GetOrPanic(func() (string, error) {
				return "hello, world!", nil
			})

			g.Expect(val).To(gomega.Equal("hello, world!"))
		},
		"should panic given a non-nil error": func(g gomega.Gomega) {
			g.Expect(func() {
				must.GetOrPanic(func() (string, error) {
					return "", errors.New("boom")
				})
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

func TestMustOrFailTest(t *testing.T) {
	t.Parallel()

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should not fail the test given a nil error": func(t testing.TB, g gomega.Gomega) {
			mt := mockT(t)
			val := must.GetOrFailTest(mt, func() (string, error) {
				return "hello, world!", nil
			})

			g.Expect(val).To(gomega.Equal("hello, world!"))
			g.Expect(mt.fatal).To(gomega.BeEmpty())
		},
		"should fail the test given a non-nil error": func(t testing.TB, g gomega.Gomega) {
			mt := mockT(t)
			val := must.GetOrFailTest(mt, func() (string, error) {
				return "", errors.New("boom")
			})

			g.Expect(val).To(gomega.Equal(""))
			g.Expect(mt.fatal).To(gomega.Equal("unexpected error: boom"))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(t, gomega.NewWithT(t))
		})
	}
}

func TestDoOrPanic(t *testing.T) {
	t.Parallel()

	scns := map[string]func(g gomega.Gomega){
		"should not panic given a nil error": func(g gomega.Gomega) {
			must.DoOrPanic(func() error {
				return nil
			})
			// If we reach here, it didn't panic
		},
		"should panic given a non-nil error": func(g gomega.Gomega) {
			g.Expect(func() {
				must.DoOrPanic(func() error {
					return errors.New("boom")
				})
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

func TestDoOrFailTest(t *testing.T) {
	t.Parallel()

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should not fail the test given a nil error": func(t testing.TB, g gomega.Gomega) {
			mt := mockT(t)
			must.DoOrFailTest(mt, func() error {
				return nil
			})

			g.Expect(mt.fatal).To(gomega.BeEmpty())
		},
		"should fail the test given a non-nil error": func(t testing.TB, g gomega.Gomega) {
			mt := mockT(t)
			must.DoOrFailTest(mt, func() error {
				return errors.New("boom")
			})

			g.Expect(mt.fatal).To(gomega.Equal("unexpected error: boom"))
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

func mockT(t testing.TB) *tm {
	return &tm{TB: t}
}

func (t *tm) Fatalf(format string, args ...any) {
	t.fatal = fmt.Sprintf(format, args...)
}
