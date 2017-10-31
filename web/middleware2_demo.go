package web

import (
	"github.com/kataras/iris"
	"net/http"
	"fmt"
	"github.com/kataras/iris/context"
)

func Middleware2Serve() *iris.Application {
	app := iris.New()
	middle := iris.FromStd(traceMiddleware)
	app.Use(middle)
	app.Use(traceMiddleware2)
	app.Get("/", func(ctx context.Context) {
		ctx.Writef("hello world!")
	})
	app.Get("/admin", func(ctx context.Context) {
		panic("fuck")
		ctx.Writef("This is a admin page")
	})

	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		//defer func() {
		//	if err := recover(); err != nil {
		//		fmt.Printf("Some exception occur %v\n", err)
		//		w.Write([]byte("Some exception occur.\n"))
		//	}
		//}()

		next(w, r)
	})
	return app
}

func traceMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Printf("current path:%s\n", r.URL.Path)
	next(w, r)
}

func traceMiddleware2(ctx context.Context) {
	w, r := ctx.ResponseWriter(), ctx.Request()

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Some Exception occur %v: %s\n", err, r.URL.Path)
			w.WriteHeader(iris.StatusInternalServerError)
		}
	}()
	ctx.Next()
}
