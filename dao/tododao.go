package dao

import (
	"errors"
	db "go_todo/model/db"

	"gorm.io/gorm"
)

type ITodoDao interface {
	CreateTodo(db.Todo) (int, error)
	UpdateTodo(todo db.Todo) error
	GetTodo(id, userID int) (*db.Todo, error)
	GetAllTodos(userID int) (*[]db.Todo, error)
	GetAllTodosByState(done int8, userID int) (*[]db.Todo, error)
	DeleteTodo(id, userID int) error
}

type TodoDao struct {
	DB *gorm.DB
}

func NewTodoDao(db *gorm.DB) TodoDao {
	return TodoDao{DB: db}
}

func (td TodoDao) CreateTodo(todo db.Todo) (int, error) {
	db := td.DB.Create(&todo)
	return todo.ID, db.Error
}

func (t TodoDao) UpdateTodo(todo db.Todo) error {
	result := t.DB.Model(todo).Where(db.Todo{
		ID:     todo.ID,
		UserID: todo.UserID,
	}).Updates(&todo)

	if result.RowsAffected == 0 {
		return errors.New("no matching data found to update")
	}

	return result.Error
}

func (t TodoDao) GetTodo(id, userID int) (*db.Todo, error) {
	todo := &db.Todo{
		UserID: userID,
	}
	result := t.DB.Where(todo).First(todo, id)
	return todo, result.Error
}

func (t TodoDao) GetAllTodos(userID int) (*[]db.Todo, error) {
	todos := &[]db.Todo{}
	result := t.DB.Where(&db.Todo{UserID: userID}).Find(todos)
	return todos, result.Error
}

func (t TodoDao) GetAllTodosByState(done int8, userID int) (*[]db.Todo, error) {
	if done != 1 && done != 0 {
		return nil, errors.New("invalid value for done,should be 0(not done) or 1 (done)")
	}
	todos := &[]db.Todo{}
	result := t.DB.Where(&db.Todo{Done: done, UserID: userID}).Find(todos)
	return todos, result.Error
}

func (t TodoDao) DeleteTodo(id int, userID int) error {
	dbUser := db.Todo{
		ID:     id,
		UserID: userID,
	}
	result := t.DB.Where(&db.Todo{UserID: userID}).Delete(&dbUser)

	if result.RowsAffected == 0 {
		return errors.New("no data to delete, please provide a valid id")
	}

	return result.Error
}
