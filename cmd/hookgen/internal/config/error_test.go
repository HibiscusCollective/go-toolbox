package config_test

import (
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
	}

	for name, test := range scns {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			test(gomega.NewWithT(t))
		})
	}
}
