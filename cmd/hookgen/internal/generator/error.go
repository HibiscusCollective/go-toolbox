package generator

// ParameterErrors is a map of fields and their errors
type ParameterErrors map[string]error

// InvalidParametersError is an error that contains a map of fields and their errors
type InvalidParametersError interface {
	error
	Parameters() ParameterErrors
}

type invalidParametersError ParameterErrors

// IntoError converts a ParameterErrors into an InvalidParametersError
func (f ParameterErrors) IntoError() InvalidParametersError {
	if f == nil || len(f) == 0 {
		return nil
	}

	return invalidParametersError(f)
}

// Error returns the error message
func (f invalidParametersError) Error() string {
	return "invalid parameter(s)"
}

// Parameters gets the parameters that are in an invalid state
func (f invalidParametersError) Parameters() ParameterErrors {
	return ParameterErrors(f)
}
