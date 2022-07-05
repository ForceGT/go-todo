package config

const (
	PortNumber             = "3306"
	SqlString              = "root:password@tcp(127.0.0.1:3306)/go_todos?charset=utf8mb4&parseTime=True&loc=Local"
	Secret                 = "USER"
	RefreshSecret          = "RefreshUser"
	DurationMinutes        = 10
	RefreshDurationMinutes = 24 * 60
	Algo                   = "HS256"
)
