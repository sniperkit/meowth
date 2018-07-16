package database

import (
	"github.com/kataras/iris"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func PgsqlInit() *gorm.DB {
	host := ""
	db, err := gorm.Open("postgres", host)
	if err != nil {
		panic("failed to connect database")
	}

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() { db.Close() })
	return db
}
