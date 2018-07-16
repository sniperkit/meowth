package controller

import (
	"github.com/kataras/iris"
)

func (c *Controller) Index(ctx iris.Context) {
	ctx.Gzip(true)
	ctx.ViewData("hello", ctx.Translate("hi","Ferdi"))
	language := ctx.Values().GetString(ctx.Application().ConfigurationReadOnly().GetTranslateLanguageContextKey())
	println(language)
	if err := ctx.View("demo.html"); err != nil {
		ctx.Application().Logger().Infof(err.Error())
	}
}