package GoLife

import "context"

import (
	"github.com/marcelib/iris"
	"github.com/marcelib/iris/context"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("../app", ".html"))

	app.Get("/", func(ctx context.Context) {
		ctx.ViewData("message", "Hello world!")
		ctx.View("hello.html")
	})

	app.Run(iris.Addr(":8080"))
}
