// Package must provides a helper function to handle error by panicking
package must

import (
	"testing"
)

// GetOrPanic panics if the error is not nil
func GetOrPanic[T any](fn func() (T, error)) T {
	val, err := fn()
	if err != nil {
		panic(err)
	}
	return val
}

// GetOrFailTest fails the test if the error is not nil
func GetOrFailTest[T any](t testing.TB, fn func() (T, error)) T {
	val, err := fn()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	return val
}
