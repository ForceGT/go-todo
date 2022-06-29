package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          int `gorm:"primaryKey;autoIncrement"`
	UserID      int
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Title       string
	Description string
	DueDate     time.Time
	Done        int8
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
