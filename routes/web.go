package routes

import (
	"github.com/weeq/meowth/bootstrap"
	"github.com/weeq/meowth/app/controller"
)

func Web(app *bootstrap.Bootstrapper,ctrl *controller.Controller) {
	app.Get("/",ctrl.Index)
}
