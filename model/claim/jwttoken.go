package model

import "github.com/golang-jwt/jwt/v4"

type JWTTokenClaim struct {
	UserID   int    `json:"id" validate:"required"`
	RoleID   int    `json:"role" validate:"required"`
	Username string `json:"username" validate:"required"`
	jwt.RegisteredClaims
}
