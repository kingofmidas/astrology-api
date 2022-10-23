package helpers

import (
	"encoding/json"
	"net/http"

	errs "github.com/kingofmidas/astrology-api/internal/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Response struct {
	StatusCode int         `json:"-"`
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (r Response) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	json.NewEncoder(w).Encode(r)
}

func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	response := Response{
		Success:    true,
		StatusCode: statusCode,
		Data:       data,
	}
	response.JSON(w)
}

func NewServerErrorResponse(err error) Response {
	return Response{
		Success:    false,
		StatusCode: http.StatusInternalServerError,
	}
}

func NewAppErrorResponse(err *errs.AppError) Response {
	statusCode := errs.GetHttpStatusCode(err.Status)

	return Response{
		Success:    false,
		StatusCode: statusCode,
		Message:    err.Err,
	}
}

func Fail(w http.ResponseWriter, err error) {
	appErr, ok := err.(*errs.AppError)
	if ok && errs.GetErrorStatus(appErr) != errs.ServerError {
		NewAppErrorResponse(appErr).JSON(w)
		return
	}

	logrus.Error(err)
	NewServerErrorResponse(err).JSON(w)
}
