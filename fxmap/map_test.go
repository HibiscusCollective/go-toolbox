package fxmap_test

import (
	"testing"

	"github.com/HibiscusCollective/go-toolbox/fxmap"
	"github.com/onsi/gomega"
)

type test func(t gomega.Gomega)

func TestInvert(t *testing.T) {
	t.Parallel()

	tests := map[string]test{
		"should return nil given nil": func(g gomega.Gomega) {
			got := fxmap.Invert[int, int](nil)

			g.Expect(got).To(gomega.BeNil())
		},
		"should return empty map given empty map": func(g gomega.Gomega) {
			got := fxmap.Invert(map[int]int{})

			g.Expect(got).To(gomega.BeEmpty())
		},
		"should return a map with the values in place of the keys": func(g gomega.Gomega) {
			got := fxmap.Invert(map[int]string{1: "one", 2: "two"})

			g.Expect(got).To(gomega.Equal(map[string]int{"one": 1, "two": 2}))
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) { test(gomega.NewWithT(t)) })
	}
}
