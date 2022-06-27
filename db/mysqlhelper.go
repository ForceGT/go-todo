package db

import (
	"go_todo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {

	sqlDialector := mysql.Open(config.SqlString)
	return gorm.Open(sqlDialector, &gorm.Config{})
}
