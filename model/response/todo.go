package model

import "time"

type TodoResponseModel struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Done        int8      `json:"done"`
}
