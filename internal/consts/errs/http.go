package errs

import (
	"errors"
	"net/http"
)

type Error struct {
	Type Type
	Err  error
}

func (e *Error) Error() string {
	return e.Err.Error()
}

func (e *Error) Unwrap() error {
	return e.Err
}

func New(t Type, e error) *Error {
	return &Error{
		Type: t,
		Err:  e,
	}
}

func HTTPError(err error) (int, string) {
	var appErr *Error

	if errors.As(err, &appErr) {
		return StatusCode(appErr.Type), appErr.Error()
	}
	return http.StatusInternalServerError, "something went wrong"
}

func StatusCode(t Type) int {
	switch t {
	case BadRequest:
		return http.StatusBadRequest
	case Unauthorized:
		return http.StatusUnauthorized
	case Forbidden:
		return http.StatusForbidden
	case NotFound:
		return http.StatusNotFound
	case RequestTimeout:
		return http.StatusRequestTimeout
	case InternalServerError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
