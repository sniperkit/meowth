package database

import (
	"github.com/kataras/iris"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func MysqlInit() *gorm.DB {
	host := ""
	db, err := gorm.Open("mysql", host)
	if err != nil {
		panic("failed to connect database")
	}

	db.Set("gorm:table_options", "engine=InnoDB")
	db.Set("gorm:table_options","charset=utf8mb4")
	db.Set("gorm:table_options","collation=utf8mb4_unicode_ci")
	db.Set("gorm:table_options","strict=1")
	db.Set("gorm:table_options","parseTime=1")
	db.Set("gorm:table_options","loc=Local")

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() { db.Close() })

	return db
}
