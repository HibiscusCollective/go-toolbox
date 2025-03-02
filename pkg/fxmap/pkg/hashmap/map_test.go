package hashmap_test

import (
	"errors"
	"hash"
	"hash/crc32"
	"testing"

	"github.com/HibiscusCollective/go-toolbox/fxmap/pkg/hashmap"
	"github.com/onsi/gomega"
)

type test func(t gomega.Gomega)

func TestNew(t *testing.T) {
	t.Parallel()

	tests := map[string]test{
		"should return an error if the hash function fails": func(g gomega.Gomega) {
			_, err := hashmap.NewWith32bitHashFunction(
				map[int32]string{1: "one"},
				newBrokenHash32Fn("boom"),
			)

			var herr hashmap.HashError
			g.Expect(err).To(gomega.MatchError(func(err error) bool {
				return errors.As(err, &herr)
			}, "hashmap.HashError"))

			g.Expect(herr.HashFn().Algorithm).To(gomega.Equal("fxmap_test.brokenHash32Fn"))
			g.Expect(herr.HashFn().Size).To(gomega.Equal(32))
			g.Expect(herr.HashFn().BlockSize).To(gomega.Equal(4))
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}

func TestInvert(t *testing.T) {
	t.Parallel()

	tests := map[string]test{
		"should return nil given nil": func(g gomega.Gomega) {
			got := hashmap.Invert[int, int](nil)

			g.Expect(got).To(gomega.BeNil())
		},
		"should return empty map given empty map": func(g gomega.Gomega) {
			got := hashmap.Invert(map[int]int{})

			g.Expect(got).To(gomega.BeEmpty())
		},
		"should return a map with the values in place of the keys": func(g gomega.Gomega) {
			got := hashmap.Invert(map[int]string{1: "one", 2: "two"})

			g.Expect(got).Should(gomega.Equal(map[string]int{"one": 1, "two": 2}))
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}

type brokenHash32Fn struct {
	hash.Hash32
	err error
}

func newBrokenHash32Fn(err string) brokenHash32Fn {
	return brokenHash32Fn{
		Hash32: crc32.New(crc32.IEEETable),
		err:    errors.New(err),
	}
}

func (h brokenHash32Fn) Write(data []byte) (int, error) {
	return 0, h.err
}
