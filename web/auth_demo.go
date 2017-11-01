package web

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/basicauth"
	"time"
	"github.com/kataras/iris/context"
)

func IrisServeWithBasic() *iris.Application {
	app := iris.New()
	config := basicauth.Config{
		Users:   map[string]string{"benjamin": "12345", "wendy": "123456"},
		Realm:   "Authorization required",
		Expires: time.Duration(30) * time.Minute,
	}
	authorization := basicauth.New(config)

	app.Get("/", func(ctx context.Context) {
		ctx.Redirect("/admin")
	})

	authNeed := app.Party("/admin", authorization)
	{
		authNeed.Get("/", h)
		authNeed.Get("/profile", h)
		authNeed.Get("/settings", h)
	}
	return app
}

func h(ctx context.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	ctx.Writef("%s %s: %s", ctx.Path(), username, password)
}
