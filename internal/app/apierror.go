package app

import "net/http"

type APIError interface {
	APIError() (int, string)
	GetStatus() int
	GetMessage() string
	Error() string
	Unwrap() error
}

type sentinelAPIError struct {
	status       int
	message      string
	wrappedError error
}

func (e sentinelAPIError) Error() string {
	return e.message
}

func (e sentinelAPIError) Unwrap() error {
	return e.wrappedError
}

func (e sentinelAPIError) APIError() (int, string) {
	return e.status, e.message
}

func (e sentinelAPIError) GetStatus() int {
	return e.status
}

func (e sentinelAPIError) GetMessage() string {
	return e.message
}

func (e sentinelAPIError) Is(target error) bool {
	t, ok := target.(*sentinelAPIError)
	if !ok {
		return false
	}

	return t.status == e.status
}

func NewAPIError(status int, message string, wrappedError error) APIError {
	return sentinelAPIError{
		status:       status,
		message:      message,
		wrappedError: wrappedError,
	}
}

var (
	ErrNotFound   = &sentinelAPIError{status: http.StatusNotFound, message: "Resource not found", wrappedError: nil}
	ErrInternal   = &sentinelAPIError{status: http.StatusInternalServerError, message: "Unexpected error found", wrappedError: nil}
	ErrBadRequest = &sentinelAPIError{status: http.StatusBadRequest, message: "Invalid data", wrappedError: nil}
)
