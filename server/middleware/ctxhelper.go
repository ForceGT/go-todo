package middleware

import "github.com/labstack/echo/v4"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	RoleID   int    `json:"roleId"`
}

func GetUserFromContext(c echo.Context) User {
	return c.Get("user").(User)
}
