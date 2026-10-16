package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//方法1：使用map
		// data := map[string]interface{}{
		// 	"name":    "小王子",
		// 	"message": "hello world!",
		// 	"age":     18,
		// }

		//gin.H==map[string]interface{}

		Data := gin.H{
			"name":    "小王子",
			"message": "hello world!",
			"age":     18,
		}

		c.JSON(http.StatusOK, Data)
	})
	type msg struct {
		Name    string
		Age     int
		Message string
	}
	Msg := msg{
		Name:    "Liu",
		Age:     22,
		Message: "12334",
	}
	r.GET("./another", func(c *gin.Context) {
		data := Msg
		c.JSON(http.StatusOK, data)
	})
	r.Run()
}
