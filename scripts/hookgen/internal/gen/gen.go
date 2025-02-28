package gen

import (
	"io"
)

type Executer interface {
	Execute(io.Writer, any) error
}

func Generate(w io.Writer, tmpl Executer) error {
	// TODO
	tmpl.Execute(w, nil)

	return nil
}
