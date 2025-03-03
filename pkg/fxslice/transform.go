package fxslice

// Transform transforms a slice of type T to a slice of type R
func Transform[T, R any](src []T, f func(T) R) []R {
	xf := make([]R, len(src))

	for i, v := range src {
		xf[i] = f(v)
	}

	return xf
}
