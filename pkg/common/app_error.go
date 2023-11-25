package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewBadRequestResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewNotFoundResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewInternalServerErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewUnauthorized(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewForbidden(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusForbidden,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewBadRequestResponse(root, msg, root.Error(), key)
	}

	return NewBadRequestResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewInternalServerErrorResponse(err, "Something went wrong with DB",
		err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewBadRequestResponse(err, "Invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInsufficientBalance(err error) *AppError {
	return NewBadRequestResponse(err, "Insufficient balance", err.Error(), "ErrInsufficientBalance")
}

func ErrMapping(err error) *AppError {
	return NewInternalServerErrorResponse(err, "Error mapping", err.Error(), "ErrMapping")
}

func ErrInternal(err error) *AppError {
	return NewInternalServerErrorResponse(err, "Something went wrong in the server",
		err.Error(), "ErrInternal")
}

func ErrCannotListEntity(entity string, err error) *AppError {
	return NewInternalServerErrorResponse(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewInternalServerErrorResponse(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewInternalServerErrorResponse(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewInternalServerErrorResponse(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		err.Error(),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewNotFoundResponse(err, fmt.Sprintf("%s not found", entity),
		err.Error(), fmt.Sprintf("Err%sNotFound", entity),
	)
}

func ErrEntityAlreadyExisted(entity string, err error) *AppError {
	return NewBadRequestResponse(err, fmt.Sprintf("%s already existed", strings.ToLower(entity)),
		err.Error(), fmt.Sprintf("Err%sAlreadyExisted", entity))
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewInternalServerErrorResponse(err, fmt.Sprintf("Cannot create %s", strings.ToLower(entity)),
		err.Error(), fmt.Sprintf("ErrCannotCreate%s", entity))
}

func ErrNoPermission(err error) *AppError {
	return NewForbidden(err, fmt.Sprintf("You have no permission"),
		err.Error(), fmt.Sprintf("ErrNoPermission"))
}

func ErrEntityBlocked(entity string, err error) *AppError {
	return NewForbidden(err, fmt.Sprintf("%s is blocked", entity),
		err.Error(), fmt.Sprintf("Err%sBlocked", entity))
}

func ErrUnauthorized(err error) *AppError {
	return NewUnauthorized(err,
		fmt.Sprintf("Access is denied due to invaid credentials"),
		err.Error(), fmt.Sprintf("ErrUnauthorized"))
}
