package gen

const (
	ErrParseProjectConfigMsg = errParseProjectConfigMsg
	ErrExecuteTemplateMsg    = errExecuteTemplateMsg
)

func WithTemplate(tmpl Executer) func(*Generator) {
	return func(g *Generator) {
		g.tmpl = tmpl
	}
}
