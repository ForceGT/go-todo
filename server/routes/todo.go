package routes

import (
	reqModel "go_todo/model/request"
	resModel "go_todo/model/response"
	"go_todo/server/controller"
	"go_todo/server/middleware"
	"go_todo/server/validator"

	"github.com/labstack/echo/v4"
)

func Todo(g *echo.Group, controller controller.ITodoController) {

	e := g.Group("todo")

	e.GET("", func(c echo.Context) error {
		user := middleware.GetUserFromContext(c)
		todoReq := &reqModel.CreateTodoRequest{}
		bindErr := validator.BindAndValidateWith(c, todoReq, validator.BindBody)
		if bindErr != nil {
			return resModel.BadRequest(bindErr)
		}
		queryResult, err := controller.CreateTodo(*todoReq, user.ID)
		if err != nil {
			return resModel.InternalServerError(err)
		}
		return resModel.SuccessJSON(c, queryResult)
	})
}
