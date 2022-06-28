package model

type CreateRoleRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=25"`
	AccessLevel int    `json:"accessLevel" validate:"required"`
}
