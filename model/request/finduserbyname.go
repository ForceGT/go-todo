package model

type FindUserbyNameRequest struct {
	Name string `param:"username" validate:"required"`
}
