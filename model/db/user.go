package model

import "time"

type User struct {
	ID        int `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
	Username  string `gorm:"unique"`
	Password  string
	Email     string
	Mobile    string
	Token     string
	RoleID    int
	Role      Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
