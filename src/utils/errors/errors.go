package errors

import (
	"errors"
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(msg string) *RestError {
	return &RestError{
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(msg string) *RestError {
	return &RestError{
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerError(msg string) *RestError {
	return &RestError{
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
