package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的query string 参数
		name := c.Query("query")
		age := c.Query("age")
		// name := c.DefaultQuery("query", "666")
		// name, ok := c.GetQuery("query") //url里面有”query=“，则ok为true，否则为false
		// if !ok {
		// 	name = "somebody"
		// }
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run("127.0.0.1:9090")
}
