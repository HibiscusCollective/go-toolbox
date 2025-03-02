package config_test

import (
	"errors"
	"testing"

	"github.com/onsi/gomega"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
)

func TestInvalidFieldsError(t *testing.T) {
	t.Parallel()

	scns := map[string]func(g gomega.Gomega){
		"should be nil if the fields are empty": func(g gomega.Gomega) {
			err := config.FieldErrors{}.IntoError()

			g.Expect(err).To(gomega.BeNil())
		},
		"should return an error if the fields are not empty": func(g gomega.Gomega) {
			err := config.FieldErrors{
				"Name": errors.New("name field must not be empty"),
			}.IntoError()

			g.Expect(err).To(gomega.Not(gomega.BeNil()))
			g.Expect(err.Error()).To(gomega.Equal("invalid field(s)"))
			g.Expect(err.Fields()).To(gomega.Equal(config.FieldErrors{
				"Name": errors.New("name field must not be empty"),
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
