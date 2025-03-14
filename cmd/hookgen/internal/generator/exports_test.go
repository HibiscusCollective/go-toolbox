package generator

import (
	"testing"

	"github.com/HibiscusCollective/go-toolbox/pkg/must"
)

type ParameterErrorImpl = parameterError
type TemplateErrorImpl = templateError

func MustCreate(t testing.TB, fsc FSCreator, engine TemplateEngine) TemplateGenerator {
	t.Helper()

	return must.GetOrFailTest(Create(fsc, engine))(t)
}
