package generator_test

import (
	"errors"
	"testing"

	"github.com/onsi/gomega"

	generator "github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/generator"
)

func TestInvalidParametersError(t *testing.T) {
	t.Parallel()

	scns := map[string]func(g gomega.Gomega){
		"should be nil if the parameters are empty": func(g gomega.Gomega) {
			err := generator.ParameterErrors{}.IntoError()

			g.Expect(err).To(gomega.BeNil())
		},
		"should return an error if the parameters are not empty": func(g gomega.Gomega) {
			err := generator.ParameterErrors{
				"Name": errors.New("name parameter is required"),
			}.IntoError()

			g.Expect(err).To(gomega.Not(gomega.BeNil()))
			g.Expect(err.Error()).To(gomega.Equal("invalid parameter(s)"))
			g.Expect(err.Parameters()).To(gomega.Equal(generator.ParameterErrors{
				"Name": errors.New("name parameter is required"),
			}))
		},
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}
