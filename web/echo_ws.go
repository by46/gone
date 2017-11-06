package web

import (
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
	"github.com/labstack/echo/middleware"
	"fmt"
	"github.com/by46/gone/basic"
	"time"
)

func WsHandler(ctx echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		me := &basic.Person{
			Name: "benjamin",
			Age:  21,
		}
		go func(ws1 *websocket.Conn) {
			for i := 0; i <= 10; i++ {
				me := &basic.Person{
					Name:  "benjamin",
					Age:   21,
					Stamp: time.Now(),
				}
				websocket.JSON.Send(ws1, me)
				time.Sleep(2 * time.Second)
			}
		}(ws)
		for {
			err := websocket.Message.Send(ws, "hello world")
			if err != nil {
				ctx.Logger().Error(err)
			}
			err = websocket.JSON.Send(ws, me)
			if err != nil {
				ctx.Logger().Error(err)
			}
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				ctx.Logger().Error(err)
			}
			fmt.Printf("Receive message: %s\n", msg)
		}

	}).ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}

func EchoWSServe() *echo.Echo {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.GET("/echo", WsHandler)
	return app
}
