package failure

import (
	"net/http"
)

// Failure is a wrapper for error messages and codes using standard HTTP response codes.
type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error returns the error code and message in a formatted string.
func (e *Failure) Error() string {
	return e.Message
}

// GetCode returns the error code of an error interface.
func GetCode(err error) int {
	if f, ok := err.(*Failure); ok {
		return f.Code
	}
	return http.StatusInternalServerError
}

// InternalError returns a new Failure with code for internal error and message derived from an error interface.
func InternalError(err error) error {
	if err != nil {
		return &Failure{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return nil
}

// BadRequest returns a new Failure with code for bad requests.
func BadRequest(err error) error {
	if err != nil {
		return &Failure{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}
	return nil
}
