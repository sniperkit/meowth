package main

import (
	"github.com/weeq/meowth/bootstrap"
	"github.com/weeq/meowth/routes"
	"github.com/weeq/meowth/app/controller"
	"github.com/weeq/meowth/lib"
)

func App() *bootstrap.Bootstrapper {

	boot := *bootstrap.Boot()

	app := boot.NewApp()

	ctrl := &controller.Controller{
		Db:      boot.Db,
		MGO:     boot.MGO,
		Session: boot.Sessions,
	}

	api_ctrl := &controller.ApiController{ctrl}
	routes.Api(&boot, api_ctrl)
	routes.Web(&boot, ctrl)

	return app
}

func main() {
	println(lib.Getenv("MGO_USE", "True") != "True")

	app := App()

	app.Listen()
}
