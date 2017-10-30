package main

import (
	"github.com/by46/gone/web"
	"github.com/kataras/iris"
)

func main() {
	//web.IrisServe()
	//web.IrisServeWithBasic().Run(iris.Addr(":8080"))
	//web.CacheServe().Run(iris.Addr(":8080"))
	web.MiddleWareServe().Run(iris.Addr(":8080"))
}
