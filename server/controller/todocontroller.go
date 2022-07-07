package controller

import (
	"go_todo/dao"
	db "go_todo/model/db"
	reqModel "go_todo/model/request"
	resModel "go_todo/model/response"
)

type ITodoController interface {
	CreateTodo(createReq reqModel.CreateTodoRequest, userid int) (interface{}, error)
	UpdateTodo(todo reqModel.UpdateTodoRequest, userID int) error
	GetTodo(id int, userID int) (*resModel.TodoResponseModel, error)
	GetAllTodos(userID int) (*[]resModel.TodoResponseModel, error)
	GetAllTodosByState(done int8, userID int) (*[]resModel.TodoResponseModel, error)
	DeleteTodo(id, userID int) error
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

func (t TodoController) UpdateTodo(todo reqModel.UpdateTodoRequest, userID int) error {
	dbTodo := db.Todo{
		ID:          todo.ID,
		UserID:      userID,
		Title:       todo.Title,
		Description: todo.Description,
		DueDate:     todo.DueDate,
		Done:        todo.Done,
	}

	return t.Dao.UpdateTodo(dbTodo)
}

func (t TodoController) GetTodo(id int, userID int) (*resModel.TodoResponseModel, error) {
	dbTodo, err := t.Dao.GetTodo(id, userID)
	if err != nil {
		return nil, err
	}

	resTodo := &resModel.TodoResponseModel{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: dbTodo.Description,
		DueDate:     dbTodo.DueDate,
		Done:        dbTodo.Done,
	}

	return resTodo, nil
}

func (t TodoController) GetAllTodos(userID int) (*[]resModel.TodoResponseModel, error) {

	todos, err := t.Dao.GetAllTodos(userID)
	if err != nil {
		return nil, err
	}

	resTodos := getResTodos(todos)
	return resTodos, nil
}
func (t TodoController) GetAllTodosByState(done int8, userID int) (*[]resModel.TodoResponseModel, error) {
	todos, err := t.Dao.GetAllTodosByState(done, userID)
	if err != nil {
		return nil, err
	}

	return getResTodos(todos), nil
}

func (t TodoController) DeleteTodo(id int, userID int) error {

	return t.Dao.DeleteTodo(id, userID)
}

func getResTodos(todos *[]db.Todo) *[]resModel.TodoResponseModel {
	var resTodos = []resModel.TodoResponseModel{}
	for _, todo := range *todos {
		resTodo := resModel.TodoResponseModel{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			DueDate:     todo.DueDate,
			Done:        todo.Done,
		}
		resTodos = append(resTodos, resTodo)
	}

	return &resTodos
}
