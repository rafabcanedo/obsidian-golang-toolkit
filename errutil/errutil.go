package errutil

import "net/http"

// AppError represents an application-level error with an associated HTTP status code.
type AppError struct {
	Message string
	StatusCode int
	Err error
}

// Error implements the error interface.
func (e *AppError) Error() string {
	return e.Message
}

// Unwrap allows errors.Is and errors.As to access the wrapped error.
func (e *AppError) Unwrap() error {
	return e.Err
}

// NotFound creates an AppError with HTTP 404 status.
func NotFound(message string) *AppError {
	return &AppError{
		Message: message,
		StatusCode: http.StatusNotFound,
	}
}

// BadRequest creates an AppError with HTTP 400 status.
func BadRequest(message string) *AppError {
	return &AppError{
		Message: message,
		StatusCode: http.StatusBadRequest,
	}
}

// Unauthorized creates an AppError with HTTP 401 status.
func Unauthorized(message string) *AppError {
	return &AppError{
		Message: message,
		StatusCode: http.StatusUnauthorized,
	}
}
