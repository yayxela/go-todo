package values

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	code    int
	message string
}

func NewHTTPError(code int, message string) *HTTPError {
	return &HTTPError{
		code:    code,
		message: message,
	}
}

func (h *HTTPError) Error() string {
	return fmt.Sprintf("http err. code:%d msg:%s", h.code, h.message)
}

func (h *HTTPError) Code() int {
	return h.code
}

func (h *HTTPError) Message() string {
	return h.message
}

var (
	ExistsError = NewHTTPError(http.StatusNotFound, "api.exists_error")
)
