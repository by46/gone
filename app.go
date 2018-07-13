package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	app := echo.New()
	app.GET("/home", func(context echo.Context) error {
		return context.HTML(http.StatusOK, "<html></html>")
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app,
	}
	srv.ListenAndServe()
}
