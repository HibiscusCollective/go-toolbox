package fxmap

func Invert[K, V comparable](m map[K]V) map[V]K {
	if m == nil {
		return nil
	}

	r := make(map[V]K, len(m))
	for k, v := range m {
		r[v] = k
	}

	return r
}
