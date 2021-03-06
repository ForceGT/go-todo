package controller

import (
	"go_todo/dao"
	db "go_todo/model/db"
	reqModel "go_todo/model/request"
	resModel "go_todo/model/response"
)

type IUserController interface {
	CreateUser(createUserRequest reqModel.CreateUserRequest) (int, error)
	FindUserByUsername(findUserReq reqModel.FindUserbyNameRequest) (resModel.User, error)
	FindUserByUserId(findUserReq reqModel.FindUserByIdRequest) (resModel.User, error)
	FindDBUserByUsername(username string) (db.User, error)
	FindDBUserByUserID(userId int) (db.User, error)
	DeleteUser(id int) error
	UpdateUserToken(*db.User) error
}

type UserController struct {
	dao dao.IUserDao
}

func NewUserController(dao dao.IUserDao) *UserController {
	return &UserController{dao: dao}
}

func (uc UserController) CreateUser(createUserRequest reqModel.CreateUserRequest) (int, error) {
	dbUser := db.User{
		FirstName: createUserRequest.FirstName,
		LastName:  createUserRequest.LastName,
		Username:  createUserRequest.Username,
		Password:  createUserRequest.Password,
		Email:     createUserRequest.Email,
		Mobile:    createUserRequest.Mobile,
		RoleID:    createUserRequest.RoleID,
	}
	return uc.dao.CreateUser(dbUser)
}

func (uc UserController) FindUserByUsername(findUserReq reqModel.FindUserbyNameRequest) (resModel.User, error) {
	value, err := uc.dao.FindUserByName(findUserReq.Name)
	resUserModel := resModel.User{
		ID:        value.ID,
		FirstName: value.FirstName,
		LastName:  value.LastName,
		Username:  value.Username,
		Email:     value.Email,
		Mobile:    value.Mobile,
		RoleID:    value.RoleID,
	}
	return resUserModel, err
}

func (uc UserController) FindUserByUserId(findUserReq reqModel.FindUserByIdRequest) (resModel.User, error) {
	value, err := uc.dao.FindUserByUserId(findUserReq.ID)
	resUserModel := resModel.User{
		ID:        value.ID,
		FirstName: value.FirstName,
		LastName:  value.LastName,
		Username:  value.Username,
		Email:     value.Email,
		Mobile:    value.Mobile,
		RoleID:    value.RoleID,
	}
	return resUserModel, err
}

func (uc UserController) FindDBUserByUsername(username string) (db.User, error) {
	value, err := uc.dao.FindUserByName(username)
	return value, err
}

func (uc UserController) DeleteUser(id int) error {
	err := uc.dao.DeleteUser(id)
	return err
}

func (uc UserController) UpdateUserToken(user *db.User) error {
	result := uc.dao.UpdateUserToken(user)
	return result
}

func (uc UserController) FindDBUserByUserID(userId int) (db.User, error) {
	result, err := uc.dao.FindUserByUserId(userId)
	return result, err
}
