package routes

import (
	"github.com/weeq/meowth/bootstrap"
	"github.com/weeq/meowth/app/controller"
)

func Api(app *bootstrap.Bootstrapper,ctrl *controller.ApiController) {
	apiRoutes := app.Party("/api")
	apiRoutes.Get("/",ctrl.Index)
}
