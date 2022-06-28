package model

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status int `json:"status"`
	Body   any `json:"body"`
}

func Success(data any) Response {
	return Response{
		Status: http.StatusOK,
		Body:   data,
	}
}

func SuccessJSON(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}
