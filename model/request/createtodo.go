package model

import "time"

type CreateTodoRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"dueDate,omitempty"`
	Done        int8      `json:"done" validate:"required,oneof=1,2"`
}
