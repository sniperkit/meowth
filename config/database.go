package config

import (
	"github.com/weeq/meowth/lib"

	"github.com/weeq/meowth/lib/database"
	"github.com/jinzhu/gorm"
)

var DefaultConnection = "sqlite"

func DatabaseInit() *gorm.DB {
	switch lib.Getenv("DB_CONNECTION",DefaultConnection) {
	case "sqlite":
		return database.SqliteInit()

	case "mysql":
		return database.MysqlInit()

	case "pgsql":
		return database.PgsqlInit()

	default:
		return database.SqliteInit()
	}
}