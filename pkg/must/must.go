// Package must provides a helper function to handle error by panicking
package must

import (
	"fmt"
	"testing"
)

// GetOrPanic panics if the error is not nil
func GetOrPanic[T any](fn func() (T, error)) T {
	val, err := fn()
	if err != nil {
		panic(fmt.Sprintf("unexpected error: %v", err))
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

// DoOrPanic panics if the error is not nil
func DoOrPanic(fn func() error) {
	if err := fn(); err != nil {
		panic(fmt.Sprintf("unexpected error: %v", err))
	}
}

// DoOrFailTest fails the test if the error is not nil
func DoOrFailTest(t testing.TB, fn func() error) {
	if err := fn(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
