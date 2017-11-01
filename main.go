package main

import (
	"github.com/by46/gone/web"
	"net/http"
)

func main() {
	//web.IrisServe()
	//web.IrisServeWithBasic().Run(iris.Addr(":8080"))
	//web.CacheServe().Run(iris.Addr(":8080"))
	//web.MiddleWareServe().Run(iris.Addr(":8080"))
	//web.Middleware2Serve().Run(iris.Addr(":8080"))
	//web.EchoServe().Start(":8080")
	//web.EchoGracefulServe()
	//web.EchoStreamServe().Start(":8080")
	//web.EchoMiddlewareServe().Start(":8080")
	//handler := web.GinServe()
	//handler := web.NativeServe()
	handler := web.GinRouterServe()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	srv.ListenAndServe()
}
