package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string   `form:"username"`
	Password string   `form:"password"`
	Age      int      `form:"age"`
	Friend   []string `form:"friend"`
}

func main() {
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		// username := c.Query("username")
		// password := c.Query("password")
		// user := UserInfo{
		// 	username: username,
		// 	password: password,
		// }

		var user UserInfo

		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
		} else {
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"method": "GET",
			})
		}
	})
	router.POST("/user", func(c *gin.Context) {
		// username := c.Query("username")
		// password := c.Query("password")
		// user := UserInfo{
		// 	username: username,
		// 	password: password,
		// }

		var user UserInfo

		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err,
			})
		} else {
			fmt.Printf("%#v\n", user)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
				"method": "POST",
			})
		}
	})

	router.Run("127.0.0.1:9090")
}
