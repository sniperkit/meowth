package controller

import (
	"github.com/kataras/iris"
	"github.com/weeq/meowth/app/models"
)

func (c *ApiController) Index(ctx iris.Context) {
	user := []models.User{}
	ctx.StatusCode(202)
	ctx.JSON(iris.Map{"data": &user, "status": 202})
}
