package model

import "github.com/golang-jwt/jwt/v4"

type JWTRefreshClaim struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	jwt.RegisteredClaims
}
