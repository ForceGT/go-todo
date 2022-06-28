package controller

import (
	"go_todo/dao"
	db "go_todo/model/db"
	requestModel "go_todo/model/request"
	responseModel "go_todo/model/response"
)

type IRoleController interface {
	CreateRole(roleRequest requestModel.CreateRoleRequest) (int, error)
	FindRoleByID(id int) (responseModel.Role, error)
}

type RoleController struct {
	Dao dao.IRoleDao
}

func NewRoleController(roleDao dao.IRoleDao) *RoleController {
	return &RoleController{
		Dao: roleDao,
	}
}

func (rc RoleController) CreateRole(roleRequest requestModel.CreateRoleRequest) (int, error) {
	dbRole := db.Role{
		Name:        roleRequest.Name,
		AccessLevel: roleRequest.AccessLevel,
	}
	return rc.Dao.CreateRole(dbRole)
}

func (rc RoleController) FindRoleByID(id int) (responseModel.Role, error) {
	queryResult, err := rc.Dao.FindRoleByID(id)
	response := responseModel.Role{
		ID:          queryResult.ID,
		Name:        queryResult.Name,
		AccessLevel: queryResult.AccessLevel,
	}
	return response, err
}
