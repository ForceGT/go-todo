package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Status      int    `json:"status"`
	ErrorString string `json:"error"`
	Message     string `json:"message"`
}

func (er ErrorResponse) Error() string {
	return er.Message
}

func BadRequest(err error) *ErrorResponse {
	return errorResponse(err, http.StatusBadRequest)
}

func UnprocessableEntity(err error) *ErrorResponse {
	return errorResponse(err, http.StatusUnprocessableEntity)
}

func InternalServerError(err error) *ErrorResponse {
	return errorResponse(err, http.StatusInternalServerError)
}

func Unauthorized(err error) *ErrorResponse {
	return errorResponse(err, http.StatusUnauthorized)
}

func NotFound(msg string) *ErrorResponse {
	return &ErrorResponse{
		Status:      http.StatusNotFound,
		ErrorString: http.StatusText(http.StatusNotFound),
		Message:     msg,
	}
}

func errorResponse(err error, httpCode int) *ErrorResponse {
	return &ErrorResponse{
		Status:      httpCode,
		ErrorString: http.StatusText(httpCode),
		Message:     err.Error(),
	}
}

func ErrorJSON(c echo.Context, e error) error {
	if err, ok := e.(ErrorResponse); ok {
		return c.JSON(err.Status, err)
	}
	return errorResponse(e, http.StatusOK)
}
