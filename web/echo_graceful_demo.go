package web

import (
	"github.com/labstack/echo"
	"time"
	"net/http"
	"github.com/tylerb/graceful"
)

func EchoGracefulServe() {
	app := echo.New()
	app.GET("/graceful", func(ctx echo.Context) error {
		time.Sleep(5 * time.Second)
		return ctx.JSON(http.StatusOK, "OK")
	})
	app.Server.Addr = ":8080"
	graceful.ListenAndServe(app.Server, 5*time.Second)
}
