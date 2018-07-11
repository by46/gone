package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	app := echo.New()

	app.GET("/", func(context echo.Context) error {
		return context.Redirect(http.StatusFound, "/")
	})
	app.GET("/close", func(context echo.Context) error {
		context.Response().Header().Set("Connection", "close")
		return context.HTML(http.StatusOK, "<html></html>")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}
	srv.ListenAndServe()
}
