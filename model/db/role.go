package model

import "time"

type Role struct {
	ID          int    `gorm:primaryKey;autoIncrement`
	Name        string `gorm:not null`
	AccessLevel int    `gorm:column:access_level;not null`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
