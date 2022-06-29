package routes

import (
	"errors"
	"go_todo/server/controller"
	"go_todo/server/validator"
	"strconv"

	"github.com/labstack/echo/v4"

	reqModel "go_todo/model/request"
	resModel "go_todo/model/response"
)

func User(g *echo.Group, controller controller.IUserController) {
	user := g.Group("/user")

	user.GET("/id/:id", func(c echo.Context) error {

		userId := &reqModel.FindUserByIdRequest{}
		bindErr := validator.BindAndValidateWith(c, userId, validator.BindPath)
		if bindErr != nil {
			return resModel.ErrorJSON(c, bindErr)
		}

		queryResultUser, err := controller.FindUserByUserId(*userId)
		if err != nil {
			return resModel.ErrorJSON(c, err)
		}

		return resModel.SuccessJSON(c, queryResultUser)
	})

	user.GET("/username/:username", func(c echo.Context) error {

		userName := &reqModel.FindUserbyNameRequest{}

		bindErr := validator.BindAndValidateWith(c, userName, validator.BindPath)

		if bindErr != nil {
			return resModel.ErrorJSON(c, bindErr)
		}

		queryResultUser, err := controller.FindUserByUsername(*userName)
		if err != nil {
			return resModel.ErrorJSON(c, err)
		}
		return resModel.SuccessJSON(c, queryResultUser)
	})

	user.POST("/", func(c echo.Context) error {
		createUserReq := &reqModel.CreateUserRequest{}
		bindErr := validator.BindAndValidateWith(c, createUserReq, validator.BindBody)
		if bindErr != nil {
			return resModel.ErrorJSON(c, bindErr)
		}
		userID, err := controller.CreateUser(*createUserReq)
		if err != nil {
			return resModel.ErrorJSON(c, err)
		}
		resStruct := struct {
			ID int `json:"id"`
		}{
			ID: userID,
		}
		return resModel.SuccessJSON(c, resStruct)
	})

	user.DELETE("/id/:id", func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return resModel.ErrorJSON(c, errors.New("ID cannot be null"))
		}

		intId, convErr := strconv.Atoi(id)
		if convErr != nil {
			return resModel.ErrorJSON(c, convErr)
		}
		err := controller.DeleteUser(intId)
		if err != nil {
			return resModel.ErrorJSON(c, err)
		}
		resStruct := struct {
			Message string `json:"message"`
		}{
			Message: "Request Processed Successfully",
		}
		return resModel.SuccessJSON(c, resStruct)
	})
}
