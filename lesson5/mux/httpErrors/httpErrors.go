package httpErrors

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
)

var (
	ContentTypeError    = errors.New("Content-Type must be `application/json`")
	CannotDecodeError   = errors.New("Can not decode error. ")
	InternalServerError = errors.New("InternalServerError")
	NotFound            = errors.New("Not Found ")
	RequestTimeoutError = errors.New("Request Timeout Error ")
)

type RestError struct {
	ErrStatus int    `json:"code,omitempty"`
	ErrError  string `json:"message,omitempty"`
}

func NewRestError(status int, err string) RestError {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
	}
}

func NewInternelServerError() RestError {
	return RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  InternalServerError.Error(),
	}
}

func ParseError(err error) RestError {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, NotFound.Error())
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, RequestTimeoutError.Error())
	default:
		return NewInternelServerError()
	}
}
