package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinServe() http.Handler {
	app := gin.Default()
	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return app
}
