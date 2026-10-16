package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:username/:age", func(c *gin.Context) {
		username := c.Param("username")
		age := c.Param("age")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
		})
	})
	r.GET("/blog/:year/:month", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			gin.H{
				"year":  ctx.Param("year"),
				"month": ctx.Param("month"),
			},
		)
	})

	r.Run("127.0.0.1:9090")
}
