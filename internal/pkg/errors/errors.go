package errs

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jackc/pgconn"
)

const (
	NotFound    = "not found"
	ServerError = "server error"
	BadRequest  = "bad request"
)

type AppError struct {
	Status string
	Err    string
}

func (err *AppError) Error() string {
	return err.Err
}

func New(status string, err string) error {
	return &AppError{
		Status: status,
		Err:    err,
	}
}

func UnexpectedError(err error, fnc string) error {
	if err == nil {
		return nil
	}
	return &AppError{
		Status: ServerError,
		Err:    fmt.Sprintf("%s: %s", fnc, err.Error()),
	}
}

func GetHttpStatusCode(status string) int {
	switch status {
	case NotFound:
		return http.StatusNotFound
	case BadRequest:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func GetErrorStatus(err error) string {
	appErr, ok := err.(*AppError)
	if ok {
		return appErr.Status
	}
	return ServerError
}

func IsUniqueConstraintError(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			return true
		}
	}
	return false
}
