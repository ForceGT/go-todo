package middleware

import (
	"errors"
	"fmt"
	"go_todo/config"
	"go_todo/server/controller"
	"regexp"

	claim "go_todo/model/claim"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	adminPathsRegex = "(\\/api\\/user.*)|(\\/api\\/role.*)"

	adminRoleID = 2
)

func GetJWTMiddleware(controller controller.IUserController) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: config.Secret,
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			token, err := jwt.ParseWithClaims(auth, &claim.JWTTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
				if token.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
				}
				return []byte(config.Secret), nil
			})

			if err != nil {
				return nil, err
			}

			if jwtClaims, ok := token.Claims.(*claim.JWTTokenClaim); ok && token.Valid {
				roleID := jwtClaims.RoleID
				path := c.Path()

				isAdminPath, err := regexp.MatchString(adminPathsRegex, path)
				if err != nil {
					panic(fmt.Sprintf("Unexpected error occured %v", err))
				}

				if isAdminPath && roleID != adminRoleID {
					return nil, errors.New("unauthorized! Only admin can make this request")
				}

				return User{
					ID:       jwtClaims.UserID,
					Username: jwtClaims.Username,
					RoleID:   jwtClaims.RoleID,
				}, nil
			}
			return nil, errors.New("could not validate token")
		},
	})
}
