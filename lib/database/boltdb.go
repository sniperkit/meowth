package database

import (
	"github.com/kataras/iris/sessions/sessiondb/boltdb"
	"github.com/weeq/meowth/lib"
	"os"
	"github.com/kataras/iris"
)

func BoltDbInit() *boltdb.Database {
	db, err := boltdb.New(lib.DatabasePath("sessions.db"), os.FileMode(0750))

	if err != nil {
		panic(err)
	}

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() { db.Close() })

	return db
}
