package models

import "fmt"

// Response represents a response structure.
type HTTPError struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	ErrorCode  string      `json:"error_code,omitempty"`
	Errors     interface{} `json:"errors,omitempty"`
}

func NewHTTPError(statusCode int, errorCode string, errorMessage string, errors interface{}) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    errorMessage,
		ErrorCode:  errorCode,
		Errors:     errors,
	}
}

// Error implements the error interface for HTTPError
func (e *HTTPError) Error() string {
	return fmt.Sprintf("Status: %d, ErrorCode: %s, Message: %s",
		e.StatusCode, e.ErrorCode, e.Message)
}
