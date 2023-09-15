// Package derrors contains domain error types to be used across services.
package derrors

type ErrorType uint

const (
	VALIDATION ErrorType = iota
	NOT_FOUND
) 

// Error defines a generic domain error
type Error struct {
	// The application-specific error code.
	Code string
	// The human-readable error message.
	Message string
	// The error type.
	Type ErrorType
}

// ValidationError should be used when client inputs are invalid.
func NewValidationError(code string, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// NewNotFoundError should be used when a requested entity is not found.
func NewNotFoundError(code string, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// This method implement errors.Error interface.
func (e *Error) Error() string {
	return e.Message
}
