package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("./login.html", "./submit.html")
	r.GET("/login", func(c *gin.Context) {

		c.HTML(http.StatusOK, "login.html", nil)
	})

	r.POST("/login", func(c *gin.Context) {
		// c.JSON(http.StatusOK, "登陆成功")
		// username, ok := c.GetPostForm("username")
		// if !ok {
		// 	username = "sb"
		// }
		// password, ok := c.GetPostForm("password")
		// if !ok {
		// 	password = "***"
		// }
		username := c.DefaultPostForm("username", "defaultname")
		password := c.DefaultPostForm("password", "defaultpassword")
		c.HTML(http.StatusOK, "submit.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run("127.0.0.1:9090")
}
