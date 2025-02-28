package gen

import (
	"embed"
	"html/template"
)

//go:embed templates
var templates embed.FS

// Templates returns the embedded templates
func Templates() *template.Template {
	return template.Must(template.ParseFS(templates, "templates/*.tmpl"))
}
