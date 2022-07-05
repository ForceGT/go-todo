package dao

import (
	db "go_todo/model/db"

	"gorm.io/gorm"
)

type ITodoDao interface {
	CreateTodo(db.Todo) (int, error)
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
