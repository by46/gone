package web

import (
	"github.com/kataras/iris"

	"fmt"
	"github.com/kataras/iris/context"
	"time"
	"github.com/kataras/iris/cache"
)

var markdownText = []byte(`
Feature
------------

- file name.
- good date
`)

func CacheServe() *iris.Application {
	app := iris.New()

	app.Get("/", cache.Handler(10*time.Second), markdownHandler)

	return app
}

func markdownHandler(ctx context.Context) {
	fmt.Printf("handler execute %d\n", time.Now().UnixNano())
	ctx.Markdown(markdownText)
}
