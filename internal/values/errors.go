package values

import (
	"fmt"
	"net/http"
)

type HttpError struct {
	code    int
	message string
}

func NewHttpError(code int, message string) *HttpError {
	return &HttpError{
		code:    code,
		message: message,
	}
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("http err. code:%d msg:%s", h.code, h.message)
}

func (h *HttpError) Code() int {
	return h.code
}

func (h *HttpError) Message() string {
	return h.message
}

var (
	ExistsError = NewHttpError(http.StatusNotFound, "api.exists_error")
)
