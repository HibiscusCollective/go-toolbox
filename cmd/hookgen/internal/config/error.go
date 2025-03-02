package config

// FieldErrors is a map of fields and their errors
type FieldErrors map[string]error

// InvalidFieldsError is an error that contains a map of fields and their errors
type InvalidFieldsError interface {
	error
	Fields() FieldErrors
}

type invalidFieldsError FieldErrors

func (f FieldErrors) IntoError() InvalidFieldsError {
	return nil
}

// Error returns the error message
func (f invalidFieldsError) Error() string {
	return "invalid field(s)"
}

// Fields gets the fields that are in an invalid state
func (f invalidFieldsError) Fields() FieldErrors {
	return FieldErrors(f)
}
