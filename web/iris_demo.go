package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"fmt"
)

func IrisServe() *iris.Application {
	app := iris.New()
	app.RegisterView(iris.HTML("./web/html", ".html"))
	app.Get("/", func(ctx context.Context) {
		ctx.ViewData("message", "hello world!")
		ctx.View("hello.html")
	})
	app.Get("/user/{id:long}", func(ctx context.Context) {
		userId, _ := ctx.Params().GetInt64("id")
		ctx.Writef("User Id %d", userId)
	})

	// MVC
	app.Controller("/city", new(UserController))
	app.Controller("/address", new(AddressController))
	app.Controller("/movies", new(MovieController))
	fmt.Printf("listen on port 8080\n")
	return app
}
