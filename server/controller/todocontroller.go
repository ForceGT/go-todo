package controller

import (
	"go_todo/dao"
	db "go_todo/model/db"
	reqModel "go_todo/model/request"
)

type ITodoController interface {
	CreateTodo(createReq reqModel.CreateTodoRequest, userid int) (interface{}, error)
}

type TodoController struct {
	Dao dao.ITodoDao
}

func NewTodoController(dao dao.ITodoDao) TodoController {
	return TodoController{
		Dao: dao,
	}
}

func (tc TodoController) CreateTodo(createReq reqModel.CreateTodoRequest, userid int) (interface{}, error) {
	dbModel := db.Todo{
		UserID:      userid,
		Title:       createReq.Title,
		Description: createReq.Description,
		DueDate:     createReq.DueDate,
		Done:        createReq.Done,
	}
	req, err := tc.Dao.CreateTodo(dbModel)

	return struct {
		ID int `json:"id"`
	}{
		req,
	}, err

}
