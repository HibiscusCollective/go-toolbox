package generator

import (
	"errors"
	"fmt"

	"github.com/HibiscusCollective/go-toolbox/cmd/hookgen/internal/config"
)

// ParameterError is an error that contains a map of fields and their errors
type ParameterError interface {
	error
	Parameter() string
}

type parameterError struct {
	Msg   string
	Label string
}

// MissingParametersError returns a list of missing parameter errors
func MissingParametersError(name string, moreNames ...string) error {
	errs := make([]error, 0, len(moreNames)+1)
	for _, name := range append([]string{name}, moreNames...) {
		if name == "" {
			continue
		}

		errs = append(errs, parameterError{
			Msg:   fmt.Sprintf("missing required parameter: %s", name),
			Label: name,
		})
	}

	if len(errs) == 0 {
		return nil
	}

	return errors.Join(errs...)
}

func (p parameterError) Error() string {
	return p.Msg
}

func (p parameterError) Parameter() string {
	return p.Label
}

// TemplateError is an error that contains the template and the data
type TemplateError interface {
	error
	Unwrap() error

	Template() string
	Data() config.Project
}

type templateError struct {
	template string
	data     config.Project
	err      error
}

// TemplateExecutionError returns an error that contains the template and the data
func TemplateExecutionError(err error, template string, data config.Project) TemplateError {
	if err == nil {
		return nil
	}

	return templateError{
		template: template,
		data:     data,
		err:      err,
	}
}

func (t templateError) Error() string {
	return "error applying template(s)"
}

func (t templateError) Unwrap() error {
	return t.err
}

func (t templateError) Template() string {
	return t.template
}

func (t templateError) Data() config.Project {
	return t.data
}
