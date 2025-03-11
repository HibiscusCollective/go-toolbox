// Package must provides a helper function to handle error by panicking
package must

import (
	"fmt"
	"testing"
)

// GetOrPanic panics if the error is not nil
func GetOrPanic[T any](val T, err error) T {
	if err != nil {
		panic(fmt.Sprintf("unexpected error: %v", err))
	}
	return val
}

// GetOrFailTest fails the test if the error is not nil
func GetOrFailTest[T any](val T, err error) func(t testing.TB) T {
	return func(t testing.TB) T {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		return val
	}
}

// DoOrPanic panics if the error is not nil
func DoOrPanic(err error) {
	if err != nil {
		panic(fmt.Sprintf("unexpected error: %v", err))
	}
}

// DoOrFailTest fails the test if the error is not nil
func DoOrFailTest(t testing.TB, err error) {
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
