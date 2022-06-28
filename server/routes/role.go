package routes

import (
	reqModel "go_todo/model/request"
	resModel "go_todo/model/response"
	"go_todo/server/controller"

	"go_todo/server/validator"

	"github.com/labstack/echo/v4"
)

func Role(g *echo.Group, controller controller.IRoleController) {
	role := g.Group("/role")
	role.POST("/", func(c echo.Context) error {
		role := &reqModel.CreateRoleRequest{}
		err := validator.BindAndValidateWith(c, role, validator.BindBody)
		if err != nil {
			return resModel.ErrorJSON(c, *err)
		}

		queryResponse, queryError := controller.CreateRole(*role)
		if queryError != nil {
			return resModel.ErrorJSON(c, queryError)
		}
		return resModel.SuccessJSON(c, struct {
			ID int `json:"id"`
		}{ID: queryResponse})
	})

	role.GET("/:id", func(c echo.Context) error {

		findRoleRequest := &reqModel.FindRoleRequest{}
		bindErr := validator.BindAndValidateWith(c, findRoleRequest, validator.BindPath)
		if bindErr != nil {
			return resModel.ErrorJSON(c, *bindErr)
		}

		resRole, err := controller.FindRoleByID(findRoleRequest.ID)
		if err != nil {
			return resModel.ErrorJSON(c, err)
		}

		return resModel.SuccessJSON(c, resRole)
	})
}
