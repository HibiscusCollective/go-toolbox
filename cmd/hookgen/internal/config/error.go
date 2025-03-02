package config

// FieldErrors is a map of fields and their errors
type FieldErrors map[string]error

// InvalidFieldsError is an error that contains a map of fields and their errors
type InvalidFieldsError interface {
	error
	Fields() FieldErrors
}

type invalidFieldsError FieldErrors

// IntoError converts a FieldErrors into an InvalidFieldsError
func (f FieldErrors) IntoError() InvalidFieldsError {
	if f == nil || len(f) == 0 {
		return nil
	}

	return invalidFieldsError(f)
}

// Error returns the error message
func (f invalidFieldsError) Error() string {
	return "invalid field(s)"
}

// Fields gets the fields that are in an invalid state
func (f invalidFieldsError) Fields() FieldErrors {
	return FieldErrors(f)
}
