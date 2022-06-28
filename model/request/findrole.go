package model

type FindRoleRequest struct {
	ID int `param:"id" validate:"required"`
}
