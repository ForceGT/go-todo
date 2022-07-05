package model

type RefreshTokenRequest struct {
	Token        string `json:"token" validate:"required"`
	RefreshToken string `json:"refreshToken" validate:"required"`
}
