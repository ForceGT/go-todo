package dao

import (
	model "go_todo/model/db"

	"gorm.io/gorm"
)

type IRoleDao interface {
	CreateRole(model.Role) (int, error)
	DeleteRoleByName(string) error
	FindRoleByID(int) (model.Role, error)
}

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{
		db: db,
	}
}

func (d RoleDao) CreateRole(role model.Role) (int, error) {
	queryResult := d.db.Create(&role)
	return role.ID, queryResult.Error
}

func (d RoleDao) DeleteRoleByName(name string) error {
	queryResult := d.db.Where("name = ?", name)
	return queryResult.Error
}

func (d RoleDao) FindRoleByID(id int) (model.Role, error) {
	role := &model.Role{ID: id}
	result := d.db.First(role)
	return *role, result.Error
}
