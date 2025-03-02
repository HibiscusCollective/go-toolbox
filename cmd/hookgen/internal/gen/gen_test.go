package gen_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/gen"
)

func TestGen(t *testing.T) {
	t.Parallel()

	const errMsg = "failed to generate hook configurations"

	scns := map[string]func(g gomega.Gomega){
		"should return an error if the config is nil": func(g gomega.Gomega) {
			err := gen.Gen(nil, nil)

			g.Expect(err).To(gomega.MatchError(fmt.Errorf("%s: %w", errMsg, gen.ParameterErrors{
				"cwd":    errors.New("cwd argument is required"),
				"config": errors.New("config argument is required"),
			}.IntoError())))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}
