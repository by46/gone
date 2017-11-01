package web

import (
	"github.com/labstack/echo"
	"net/http"
	"github.com/labstack/echo/middleware"
	"fmt"
	"strconv"
	"os"
	"io"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}

type User struct {
	ID    int    `json:"id" xml:"id" form:"id" query:"id"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func trace(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Printf("request to /users\n")
		return next(ctx)
	}
}

func EchoServe() *echo.Echo {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.RequestID())
	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: [] string{"*"},
		AllowMethods: [] string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	app.GET("/file/profile", func(ctx echo.Context) error {
		return ctx.File("1.jpg")
	})
	app.POST("/file/profile", func(ctx echo.Context) error {
		file, err := ctx.FormFile("profile")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		dst, err := os.Create(file.Filename)
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		return ctx.HTML(http.StatusOK, fmt.Sprintf("%s uploaded", file.Filename))
	})

	app.GET("/", hello)
	app.GET("/user/:id", func(ctx echo.Context) error {
		id, _ := strconv.Atoi(ctx.Param("id"))
		return ctx.JSON(http.StatusOK, &User{ID: id, Name: "benjamin", Email: "Benjamin.C.Yan"}, )
	})

	app.POST("/users", func(ctx echo.Context) error {
		user := new(User)
		if err := ctx.Bind(user); err != nil {
			return err
		}
		return ctx.JSON(http.StatusCreated, user)
	}, trace)

	g := app.Group("/admin")
	g.Use(middleware.BasicAuth(func(name string, password string, ctx echo.Context) (bool, error) {
		if name == "benjamin" && password == "123456" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/profile", hello)

	return app
}
