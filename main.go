package main

import (
	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang!",
	})
}

func main() {
	r := gin.Default() //返回默认的路由引擎

	r.GET("/hello", sayHello)

	r.Run("127.0.0.1:9090")
}