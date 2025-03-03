// Package must provides a helper function to handle error by panicking
package must

import "testing"

// Success is an interface that defines the logic for requiring the successful of a function
type Success[T any] interface {
	OrPanic() T
	OrFail(t testing.TB) T
}

type result[T any] struct {
	v   T
	err error
}

// Succeed returns a Success that resolves to the given value and error
func Succeed[T any](val T, err error) Success[T] {
	return result[T]{
		v:   val,
		err: err,
	}
}

// OrPanic panics if the error is not nil
func (r result[T]) OrPanic() T {
	if r.err != nil {
		panic(r.err)
	}

	return r.v
}

// OrFail fails the test if the error is not nil
func (r result[T]) OrFail(t testing.TB) T {
	if r.err != nil {
		t.Fatalf("unexpected error: %v", r.err)
	}

	return r.v
}
