package fxslice_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/pkg/fxslice"
)

func TestCast(t *testing.T) {
	t.Parallel()

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should return a nil slice if the slice is nil": func(t testing.TB, g gomega.Gomega) {
			rs, err := fxslice.Cast[uint8, uint](nil)

			g.Expect(err).To(gomega.BeNil())
			g.Expect(rs).To(gomega.BeNil())
		},
		"should return an empty slice if the slice is empty": func(t testing.TB, g gomega.Gomega) {
			rs, err := fxslice.Cast[uint8, uint]([]uint8{})

			g.Expect(err).To(gomega.BeNil())
			g.Expect(rs).To(gomega.BeNil())
		},
		"should return a transformed slice": func(t testing.TB, g gomega.Gomega) {
			rs, err := fxslice.Cast[uint8, uint]([]uint8{1, 2, 3})

			g.Expect(err).To(gomega.BeNil())
			g.Expect(rs).To(gomega.Equal([]uint{1, 2, 3}))
		},
		"should return all errors that occurred during the transformation": func(t testing.TB, g gomega.Gomega) {
			rs, err := fxslice.Cast[any, uint]([]any{"abc", 1, "ghi"})

			g.Expect(err).To(gomega.MatchError(errors.Join(
				fmt.Errorf("index 0: %w", fxslice.NewCastError(any("abc"), reflect.TypeOf(uint(0)))),
				fmt.Errorf("index 2: %w", fxslice.NewCastError(any("ghi"), reflect.TypeOf(uint(0)))),
			)))
			g.Expect(rs).To(gomega.Equal([]uint{1}))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(t, gomega.NewWithT(t))
		})
	}
}

func TestTransform(t *testing.T) {
	t.Parallel()

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should return a nil slice if the slice is nil": func(t testing.TB, g gomega.Gomega) {
			rs := fxslice.Transform(nil, plusTwo)
			g.Expect(rs).To(gomega.BeNil())
		},
		"should return an empty slice if the slice is empty": func(t testing.TB, g gomega.Gomega) {
			rs := fxslice.Transform([]uint8{}, plusTwo)

			g.Expect(rs).To(gomega.Equal([]uint8{}))
		},
		"should return a transformed slice": func(t testing.TB, g gomega.Gomega) {
			rs := fxslice.Transform([]uint8{1, 2, 3}, plusTwo)

			g.Expect(rs).To(gomega.Equal([]uint8{3, 4, 5}))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(t, gomega.NewWithT(t))
		})
	}
}

func TestTryTransform(t *testing.T) {
	t.Parallel()

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should return a nil slice if the slice is nil": func(t testing.TB, g gomega.Gomega) {
			rs, err := fxslice.TryTransform(nil, succeed(plusTwo))

			g.Expect(err).To(gomega.BeNil())
			g.Expect(rs).To(gomega.BeNil())
		},
		"should return a nil slice if the slice is empty": func(t testing.TB, g gomega.Gomega) {
			rs, err := fxslice.TryTransform([]uint8{}, succeed(plusTwo))

			g.Expect(err).To(gomega.BeNil())
			g.Expect(rs).To(gomega.BeNil())
		},
		"should return a transformed slice": func(t testing.TB, g gomega.Gomega) {
			rs, err := fxslice.TryTransform([]uint8{1, 2, 3}, succeed(plusTwo))

			g.Expect(err).To(gomega.BeNil())
			g.Expect(rs).To(gomega.Equal([]uint8{3, 4, 5}))
		},
		"should return all errors that occurred during the transformation": func(t testing.TB, g gomega.Gomega) {
			rs, err := fxslice.TryTransform([]uint8{1, 2, 3}, fail(plusTwo))

			g.Expect(err).ToNot(gomega.Equal(errors.Join(errors.New("failed"), errors.New("failed"), errors.New("failed"))))
			g.Expect(rs).To(gomega.BeNil())
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(t, gomega.NewWithT(t))
		})
	}
}

func plusTwo(it uint8) uint8 {
	return it + 2
}

func succeed[T, R any](fn fxslice.Transformer[T, R]) func(T) (R, error) {
	return func(t T) (R, error) {
		return fn(t), nil
	}
}

func fail[T, R any](_ fxslice.Transformer[T, R]) func(T) (R, error) {
	return func(t T) (R, error) {
		return *new(R), errors.New("failed")
	}
}
