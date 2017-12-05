package main

import (
	"net/http"

	"github.com/by46/gone/raft"
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
	//handler := web.GinRouterServe()
	//handler := basic.GameServe()
	//handler := web.EchoWSServe()
	//handler := web.EchoGorillaServe()
	//handler := web.EchoGorilla2Serve()
	handler := raft.Serve()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
	srv.ListenAndServe()
}
