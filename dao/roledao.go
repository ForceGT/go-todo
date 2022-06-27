package dao

import "gorm.io/gorm"

type RoleDao struct {
	db *gorm.DB
}

func NewRoleDao(db *gorm.DB) *RoleDao {
	return &RoleDao{
		db: db,
	}
}
