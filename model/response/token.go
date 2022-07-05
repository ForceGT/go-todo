package model

type JWTTokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
