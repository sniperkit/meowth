package database

import (
	"github.com/weeq/meowth/lib"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris"
)

func SqliteInit() *gorm.DB {
	db, err := gorm.Open("sqlite3", lib.DatabasePath("database.db"))
	if err != nil {
		panic("failed to connect database")
	}

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() { db.Close() })
	return db
}
