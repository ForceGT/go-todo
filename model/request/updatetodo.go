package model

import "time"

type UpdateTodoRequest struct {
	ID          int       `json:"id" validate:"required"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	DueDate     time.Time `json:"dueDate,omitempty"`
	Done        int8      `json:"done,omitempty" validate:"oneof=0 1"`
}
