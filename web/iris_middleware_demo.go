package web

import (
	"github.com/kataras/iris"
	"net/http"
	"github.com/kataras/iris/context"
)

func MiddleWareServe() *iris.Application {
	app := iris.New()
	middleware := iris.FromStd(redirectMiddle)
	app.Use(middleware)
	app.Get("/", func(ctx context.Context) {
		ctx.HTML("<h1>Home</h1>")
	})
	app.Get("/ok", func(ctx context.Context) {
		ctx.Writef("Hello world!")
	})
	return app
}

func redirectMiddle(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if r.URL.Path == "/ok" && r.Method == "GET" {
		w.Write([]byte("OK."))
		next(w, r)
		return
	}
	w.WriteHeader(iris.StatusBadRequest)
	w.Write([]byte("Bad request"))
}
