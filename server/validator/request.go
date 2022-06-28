package validator

import (
	resmodel "go_todo/model/response"

	"github.com/labstack/echo/v4"
)

// Binding is done in following order: 1) path params; 2) query params;
// 3) request body when using Bind. Each step COULD override previous
// step binded values. For this reason we use methods BindBody, BindQueryParams, BindPathParams.
func BindAndValidateWith(
	c echo.Context,
	model interface{},
	mainBinder func(c echo.Context, model interface{}) *resmodel.ErrorResponse,
	binders ...func(c echo.Context, model interface{}) *resmodel.ErrorResponse,
) *resmodel.ErrorResponse {

	for _, binder := range append(binders, mainBinder) {
		bindErr := binder(c, model)
		if bindErr != nil {
			return bindErr
		}
	}

	valErr := Validate(c, model)
	if valErr != nil {
		return valErr
	}

	return nil
}

func BindBody(c echo.Context, model interface{}) *resmodel.ErrorResponse {
	binder := echo.DefaultBinder{}
	bindErr := binder.BindBody(c, model)
	return checkErr(bindErr)
}

func BindQuery(c echo.Context, model interface{}) *resmodel.ErrorResponse {
	binder := echo.DefaultBinder{}
	bindErr := binder.BindQueryParams(c, model)
	return checkErr(bindErr)
}

func BindPath(c echo.Context, model interface{}) *resmodel.ErrorResponse {
	binder := echo.DefaultBinder{}
	bindErr := binder.BindPathParams(c, model)
	return checkErr(bindErr)
}

func BindHeaders(c echo.Context, model interface{}) *resmodel.ErrorResponse {
	binder := echo.DefaultBinder{}
	bindErr := binder.BindHeaders(c, model)
	return checkErr(bindErr)
}

func Validate(c echo.Context, model interface{}) *resmodel.ErrorResponse {
	bindErr := c.Validate(model)
	if bindErr != nil {
		return resmodel.BadRequest(bindErr)
	}
	return nil
}

func checkErr(err error) *resmodel.ErrorResponse {
	if err != nil {
		return resmodel.UnprocessableEntity(err)
	}
	return nil
}
