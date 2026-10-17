package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/index", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "get",
		})

	})
	router.Any("/test", func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case http.MethodGet:
			ctx.JSON(http.StatusOK, gin.H{
				"message": "any-get",
			})
		case http.MethodPost:
			ctx.JSON(http.StatusOK, gin.H{
				"message": "any-post",
			})
		}
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "你个傻逼，没有这个界面",
		})
	})

	video := router.Group("/video")
	{
		video.GET("/a", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"path": "video/a",
			})
		})
		video.GET("/b", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"path": "video/b",
			})
		})
		recommend := video.Group("/recommend")
		{
			recommend.GET("/aa", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"path": "video/recommend/aa",
				})

			})
		}
	}
	router.Run("127.0.0.1:9090")
}
