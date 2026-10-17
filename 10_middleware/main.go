package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	fmt.Println("indexHandler in...")
	name, _ := c.Get("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "indexHandler",
		"name":    name,
	})
}

func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	start := time.Now()

	c.Next()

	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out ...")
}
func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "liu") //中间件的通信
	c.Next()
	// c.Abort()
	fmt.Println("m2 out ...")

}

func main() {
	router := gin.Default()

	router.Use(m1, m2)
	router.GET("/middle", indexHandler) //等所有的中间件彻底跑完，才会把数据返回给浏览器

	router.Run("127.0.0.1:9090")
}
