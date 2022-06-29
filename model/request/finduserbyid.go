package model

type FindUserByIdRequest struct {
	ID int `param:"id" validate:"required"`
}
