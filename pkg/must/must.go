// Package must provides a helper function to handle error by panicking
package must

func OrPanic[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
