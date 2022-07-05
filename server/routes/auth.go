package routes

import (
	"fmt"
	"go_todo/config"
	"go_todo/server/controller"
	"go_todo/server/validator"

	claim "go_todo/model/claim"
	db "go_todo/model/db"
	reqModel "go_todo/model/request"
	resModel "go_todo/model/response"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func Auth(e *echo.Echo,
	userController controller.IUserController,
	authController controller.IAuthController) {

	auth := e.Group("/auth")

	auth.POST("/login", func(c echo.Context) error {
		req := &reqModel.LoginUserRequest{}
		bindErr := validator.BindAndValidateWith(c, req, validator.BindBody)
		if bindErr != nil {
			return resModel.ErrorJSON(c, bindErr)
		}

		user, err := userController.FindDBUserByUsername(req.Username)
		if err != nil {
			return resModel.ErrorJSON(c, err)
		}

		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
		if passErr != nil {
			incorrectPassError := errors.WithMessage(passErr, "Incorrect password")
			return resModel.ErrorJSON(c, resModel.Unauthorized(incorrectPassError))
		}
		jwtTokens, jwtError := generateTokensAndUpdateUser(c, user, userController, authController)
		if jwtError != nil {
			return resModel.ErrorJSON(c, resModel.InternalServerError(jwtError))
		}

		return resModel.SuccessJSON(c, jwtTokens)
	})

	auth.POST("/refresh", func(c echo.Context) error {
		reqRefresh := &reqModel.RefreshTokenRequest{}
		bindValErr := validator.BindAndValidateWith(c, reqRefresh, validator.BindBody)
		if bindValErr != nil {
			return resModel.ErrorJSON(c, bindValErr)
		}

		refreshToken, err := jwt.ParseWithClaims(
			reqRefresh.RefreshToken,
			&claim.JWTRefreshClaim{},
			func(token *jwt.Token) (interface{}, error) {
				if token.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", token.Header["alg"])
				}
				return config.Secret, nil
			})

		if err != nil {
			return resModel.ErrorJSON(c, resModel.Unauthorized(err))
		}

		if refClaims, ok := refreshToken.Claims.(*claim.JWTRefreshClaim); ok && refreshToken.Valid {
			user, err := userController.FindDBUserByUsername(refClaims.Username)
			if err != nil {
				return resModel.ErrorJSON(c, resModel.BadRequest(errors.New("Invalid username")))
			}

			if refClaims.Subject != "refresh" {
				return resModel.ErrorJSON(c, resModel.Unauthorized(errors.New("Invalid refresh token")))
			}

			if user.ID != refClaims.ID {
				return resModel.ErrorJSON(c, resModel.Unauthorized(errors.New("Invalid refresh token for user")))
			}

			if user.Token != reqRefresh.RefreshToken {
				return resModel.ErrorJSON(c, resModel.Unauthorized(errors.New("Invalid refresh token")))
			}

			jwtTokens, jwtError := generateTokensAndUpdateUser(c, user, userController, authController)
			if jwtError != nil {
				return resModel.ErrorJSON(c, jwtError)
			}

			return resModel.SuccessJSON(c, jwtTokens)
		}

		return resModel.ErrorJSON(c, resModel.Unauthorized(errors.New("Invalid refresh token")))
	})
}

func generateTokensAndUpdateUser(
	c echo.Context, user db.User,
	userController controller.IUserController,
	authController controller.IAuthController,
) (resModel.JWTTokenResponse, *resModel.ErrorResponse) {
	jwtTokens, jwtError := authController.GenerateToken(user)
	if jwtError != nil {
		return jwtTokens, resModel.InternalServerError(jwtError)
	}

	user.Token = jwtTokens.RefreshToken
	updateErr := userController.UpdateUserToken(&user)

	if updateErr != nil {
		return jwtTokens, resModel.InternalServerError(updateErr)
	}

	return jwtTokens, nil
}
