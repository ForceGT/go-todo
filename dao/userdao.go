package dao

import (
	db "go_todo/model/db"

	"gorm.io/gorm"
)

type IUserDao interface {

	//Return the user id of the created user
	CreateUser(db.User) (int, error)

	//Update the current user
	UpdateUser(db.User) error

	//Find the user by username
	FindUserByName(string) (db.User, error)

	//Find the user by the user id
	FindUserByUserId(id int) (db.User, error)

	//Delete the user by the user id
	DeleteUser(id int) error

	//Update the refresh token of the user
	UpdateUserToken(db.User) error
}

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return UserDao{
		db: db,
	}
}

func (ud UserDao) CreateUser(user db.User) (int, error) {
	gormDb := ud.db.Create(&user)
	return user.ID, gormDb.Error
}

func (ud UserDao) UpdateUser(user db.User) error {
	result := ud.db.Model(user).Updates(&user)
	return result.Error
}

func (ud UserDao) FindUserByName(name string) (db.User, error) {
	var user db.User
	result := ud.db.Where("username = ?", name).First(&user)
	return user, result.Error

}

func (ud UserDao) FindUserByUserId(id int) (db.User, error) {
	var user db.User
	result := ud.db.Where("id = ?", id).First(&user)
	return user, result.Error

}

func (ud UserDao) DeleteUser(id int) error {
	var user db.User
	result := ud.db.Where("id = ?", id).Delete(&user)
	return result.Error
}

func (u UserDao) UpdateUserToken(user *db.User) error {
	result := u.db.Model(&user).Update("token", user.Token)
	return result.Error
}
