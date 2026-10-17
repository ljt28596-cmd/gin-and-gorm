package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/index", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	router.GET("/a", func(ctx *gin.Context) {

		ctx.Request.URL.Path = "/b"
		router.HandleContext(ctx)
	})
	router.GET("/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	router.Run("127.0.0.1:9090")
}
