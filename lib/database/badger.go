package database

import (
	"github.com/kataras/iris/sessions/sessiondb/badger"
	"github.com/weeq/meowth/lib"
	"github.com/kataras/iris"
)

func BadgerDbInit() *badger.Database {
	db, err := badger.New(lib.DatabasePath("sessions"))

	if err != nil {
		panic(err)
	}

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() { db.Close() })

	return db
}