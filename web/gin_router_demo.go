package web

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
)

type Login struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		t := time.Now()
		ctx.Set("example", "123456")
		ctx.Next()
		latency := time.Since(t)
		fmt.Printf("status %v, elapse %v\n", ctx.Writer.Status(), latency)
	}
}

func GinRouterServe() http.Handler {
	app := gin.Default()
	app.Use(Logger())
	app.GET("/string/:name", func(ctx *gin.Context) {
		age := ctx.DefaultQuery("age", "16")
		name := ctx.Param("name")
		fmt.Printf("Hello %s %s\n", name, age)
	})
	app.POST("/login", func(ctx *gin.Context) {
		var json Login
		if ctx.Bind(&json) == nil {
			fmt.Printf("Username %s Login \n", json.Name)
			ctx.JSON(http.StatusOK, gin.H{"Status": true})
		}
	})
	app.GET("/login", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &Login{"Benjamin", "123545"})
	})
	return app
}
