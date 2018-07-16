package database

import (
	"github.com/weeq/meowth/lib"
	"fmt"
	"log"
	"github.com/kataras/iris"
	"gopkg.in/mgo.v2"
)

func MgoInit() *mgo.Database {

	if (lib.Getenv("MGO_USE", "True") != "True") {
		return nil
	}

	host := lib.Getenv("MGO_HOST", "127.0.0.1")
	port := lib.Getenv("MGO_PORT", "27017")
	Database := lib.Getenv("MGO_DATABASE", "ads")
	Username := lib.Getenv("MGO_USERNAME", "ads")
	Password := lib.Getenv("MGO_PASSWORD", "")
	mgoHost := fmt.Sprintf("mongodb://%v:%v@%v:%v/%v", Username, Password, host, port, Database)

	conn, err := mgo.Dial(mgoHost)

	connection := conn.Copy()

	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	// close connection when control+C/cmd+C
	iris.RegisterOnInterrupt(func() { connection.Close() })

	return connection.DB(lib.Getenv("MGO_DATABASE", ""))
}
