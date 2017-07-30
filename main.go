package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type HelloPage struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html"))

	app.Get("/", func(ctx context.Context) {
		ctx.View("endpoint.html")
	})

	app.Get("/index", func(ctx context.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.ContentType("application/javascript")
		page := HelloPage{
			Title:       "Welcome to my new project 'GoLife'!",
			Description: "This will be a multiplayer version of conway's 'Game of Life'.",
		}
		ctx.JSON(page)
	})

	app.Run(iris.Addr(":8080"))
}
