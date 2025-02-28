package gen

func WithTemplate(tmpl Executer) func(*Generator) {
	return func(g *Generator) {
		g.tmpl = tmpl
	}
}
