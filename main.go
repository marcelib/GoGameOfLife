package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html"))

	app.Get("/", func(ctx context.Context) {
		ctx.View("endpoint.html")
	})

	app.Get("/message", func(ctx context.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.JSON("Welcome to My new project, GoLife!")
	})

	app.Run(iris.Addr(":8080"))
}
