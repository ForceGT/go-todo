package model

type TodoByID struct {
	ID int `param:"id" validate:"required"`
}
