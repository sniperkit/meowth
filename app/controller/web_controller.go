package controller

import "github.com/kataras/iris"

func (c *Controller) Index(ctx iris.Context) {
	ctx.Writef("Hello World")
}