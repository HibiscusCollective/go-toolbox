package fxslice_test

import (
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/pkg/fxslice"
)

func TestTransform(t *testing.T) {
	t.Parallel()

	scns := map[string]func(t testing.TB, g gomega.Gomega){
		"should return an empty slice if the slice is empty": func(t testing.TB, g gomega.Gomega) {
			g.Expect(fxslice.Transform([]uint8{}, plusTwo)).To(gomega.Equal([]uint8{}))
		},
		"should return a transformed slice": func(t testing.TB, g gomega.Gomega) {
			g.Expect(fxslice.Transform([]uint8{1, 2, 3}, plusTwo)).To(gomega.Equal([]uint8{3, 4, 5}))
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
