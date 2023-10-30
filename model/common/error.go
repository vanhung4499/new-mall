package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Error struct {
	Code int    `json:"code"`
	Err  error  `json:"-"`
	Msg  string `json:"message"`
	Log  string `json:"log"`
	Key  string `log:"key"`
}

func NewErrorResponse(err error, msg, log, key string) *Error {
	return &Error{
		Code: http.StatusBadRequest,
		Err:  err,
		Msg:  msg,
		Log:  log,
		Key:  key,
	}
}

func NewFullErrorResponse(statusCode int, err error, msg, log, key string) *Error {
	return &Error{statusCode, err, msg, log, key}
}

func NewUnauthorized(err error, msg, key string) *Error {
	return &Error{Code: http.StatusUnauthorized, Err: err, Msg: msg, Key: key}
}

func NewCustomError(err error, msg string, key string) *Error {
	if err != nil {
		return NewErrorResponse(err, msg, err.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *Error) RootError() error {
	var err *Error
	switch {
	case errors.As(e.Err, &err):
		return err.RootError()
	default:
		return e.Err
	}
}

func (e *Error) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *Error {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *Error {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *Error {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "something went wrong in the server", err.Error(), "ErrInternal")
}
func ErrCannotListEntity(entity string, err error) *Error {
	return NewCustomError(err, fmt.Sprintf("Cannot list %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotList%s", entity))
}

func ErrCannotDeleteEntity(entity string, err error) *Error {
	return NewCustomError(err, fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotDelete%s", entity))
}

func ErrEntityDeleted(entity string, err error) *Error {
	return NewCustomError(err, fmt.Sprintf("%s deleted", strings.ToLower(entity)), fmt.Sprintf("Err%sDeleted", entity))
}

func ErrEntityExisted(entity string, err error) *Error {
	return NewCustomError(err, fmt.Sprintf("%s already exists", strings.ToLower(entity)), fmt.Sprintf("Err%sAlreadyExists", entity))
}

func ErrEntityNotFound(entity string, err error) *Error {
	return NewCustomError(err, fmt.Sprintf("%s not found", strings.ToLower(entity)), fmt.Sprintf("Err%sNotFound", entity))
}

func ErrCannotCreateEntity(entity string, err error) *Error {
	return NewCustomError(err, fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreate%s", entity))
}

func ErrNoPermission(err error) *Error {
	return NewCustomError(err, fmt.Sprintf("You have no permission"), fmt.Sprintf("ErroNoPermission"))
}

var RecordNotFound = errors.New("record not found")
