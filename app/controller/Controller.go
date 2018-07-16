package controller

import (
	"github.com/kataras/iris/sessions"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

type Controller struct {
	Db      *gorm.DB
	MGO     *mgo.Database
	Session *sessions.Sessions
}

type ApiController struct {
	*Controller
}
