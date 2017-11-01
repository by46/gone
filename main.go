package main

import (
	"github.com/by46/gone/web"
)

func main() {
	//web.IrisServe()
	//web.IrisServeWithBasic().Run(iris.Addr(":8080"))
	//web.CacheServe().Run(iris.Addr(":8080"))
	//web.MiddleWareServe().Run(iris.Addr(":8080"))
	//web.Middleware2Serve().Run(iris.Addr(":8080"))
	//web.EchoServe().Start(":8080")
	web.EchoGracefulServe()
}
