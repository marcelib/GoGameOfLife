package main
import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)
func main() {
	app := iris.New()
	// Load all templates from the "./templates" folder
	// where extension is ".html" and parse them
	// using the standard `html/template` package.
	app.RegisterView(iris.HTML("./templates", ".html"))

	// Method:    GET
	// Resource:  http://localhost:8080
	app.Get("/", func(ctx context.Context) {
		// Bind: {{.message}} with "Hello world!"
		ctx.ViewData("message", "Hello world!")
		// Render template file: ./templates/hello.html
		ctx.View("hello.html")
	})

	// Start the server using a network address and block.
	app.Run(iris.Addr(":8080"))
}