package routes

import (
	"errors"
	reqModel "go_todo/model/request"
	resModel "go_todo/model/response"
	"go_todo/server/controller"
	"go_todo/server/middleware"
	"go_todo/server/validator"
	"net/http"
	"strings"

	"strconv"

	"github.com/labstack/echo/v4"
)

func Todo(g *echo.Group, controller controller.ITodoController) {

	e := g.Group("todo")

	e.POST("", func(c echo.Context) error {
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

	e.GET("/:id", func(c echo.Context) error {
		user := middleware.GetUserFromContext(c)
		todoReq := &reqModel.TodoByID{}

		bindErr := validator.BindAndValidateWith(c, todoReq, validator.BindPath)

		if bindErr != nil {
			return resModel.BadRequest(bindErr)
		}

		todo, err := controller.GetTodo(todoReq.ID, user.ID)
		if err != nil {

			return resModel.InternalServerError(err)
		}
		return resModel.SuccessJSON(c, todo)
	})

	e.GET("/all", func(c echo.Context) error {

		user := middleware.GetUserFromContext(c)
		done := c.QueryParam("done")
		var todoResults *[]resModel.TodoResponseModel
		if done != "" {
			// We want a list based on the done status
			if (strings.Compare("1", done) != 0) || (strings.Compare("0", done) != 0) {
				return resModel.BadRequest(errors.New("done should be either 0(not done) or 1(done)"))
			}
			if doneInt, err := strconv.Atoi(done); err == nil {
				todoList, err := controller.GetAllTodosByState(int8(doneInt), user.ID)
				if err != nil {
					resModel.InternalServerError(err)
				}
				todoResults = todoList

			} else {
				return resModel.BadRequest(errors.New("done should be either 0(not done) or 1(done)"))
			}
		} else {

			todos, err := controller.GetAllTodos(user.ID)
			if err != nil {
				resModel.InternalServerError(err)
			}
			todoResults = todos
		}
		return resModel.SuccessJSON(c, todoResults)
	})

	e.PUT("", func(c echo.Context) error {
		user := middleware.GetUserFromContext(c)

		updateReq := reqModel.UpdateTodoRequest{}

		bindErr := validator.BindAndValidateWith(c, &updateReq, validator.BindBody)
		if bindErr != nil {
			return resModel.BadRequest(bindErr)
		}

		err := controller.UpdateTodo(updateReq, user.ID)
		if err != nil {
			return resModel.InternalServerError(err)
		}

		return c.NoContent(http.StatusOK)
	})

}
