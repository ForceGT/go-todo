package controller

import (
	"go_todo/config"
	claims "go_todo/model/claim"
	db "go_todo/model/db"
	resModel "go_todo/model/response"
	"time"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
)

type IAuthController interface {
	GenerateToken(user db.User) (resModel.JWTTokenResponse, error)
}

type AuthController struct {
	secret        []byte
	refreshSecret []byte
	ttl           time.Duration
	refreshTTL    time.Duration
	algo          jwt.SigningMethod
}

func NewAuthController() AuthController {
	return AuthController{
		secret:        []byte(config.Secret),
		refreshSecret: []byte(config.RefreshSecret),
		ttl:           time.Duration(config.DurationMinutes) * time.Minute,
		refreshTTL:    time.Duration(config.RefreshDurationMinutes) * time.Minute,
		algo:          jwt.GetSigningMethod(config.Algo),
	}
}

func (ac AuthController) GenerateToken(user db.User) (resModel.JWTTokenResponse, error) {
	jwtClaims := &claims.JWTTokenClaim{
		UserID:   user.ID,
		RoleID:   user.RoleID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ac.ttl)),
		},
	}

	refreshClaims := claims.JWTRefreshClaim{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ac.refreshTTL)),
			Subject:   "refresh",
		},
	}

	err := validator.New().Struct(jwtClaims)
	if err != nil {
		return resModel.JWTTokenResponse{}, nil
	}

	token, tokenErr := jwt.NewWithClaims(ac.algo, jwtClaims).SignedString(ac.secret)
	if tokenErr != nil {
		return resModel.JWTTokenResponse{}, tokenErr
	}

	refreshToken, refreshTokenErr := jwt.NewWithClaims(ac.algo, refreshClaims).SignedString(ac.refreshSecret)
	if refreshTokenErr != nil {
		return resModel.JWTTokenResponse{}, refreshTokenErr
	}

	return resModel.JWTTokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}
