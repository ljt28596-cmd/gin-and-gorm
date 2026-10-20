package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GORM 核心库

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var DB *gorm.DB

// 连接数据库
func initMySQL() (err error) {
	dsn := "root:root1234@tcp(127.0.0.1:13306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil

}
func main() {

	err := initMySQL()
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&Todo{}) //转化为数据库表时，表名为todos

	router := gin.Default()
	//找静态文件
	router.Static("/static", "static") //第一个参数：url上的路径，就是对外暴露的路径。第二个参数是电脑上的实际路径
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	v1Group := router.Group("/v1")
	{
		//添加
		v1Group.POST("/todo", func(ctx *gin.Context) {
			//从请求中把数据拿出来
			var todo Todo
			ctx.BindJSON(&todo)
			//存入数据库
			//返回数据
			if err = DB.Create(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, todo)
			}
		})
		//查看所有
		v1Group.GET("/todo", func(ctx *gin.Context) {
			//查看todos表里的所有数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, todoList)
			}
		})
		//查看单个
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {

		})

		//修改
		v1Group.PUT("/todo/:id", func(ctx *gin.Context) {
			// id, ok := ctx.Params.Get("id")
			// if !ok {
			// 	ctx.JSON(http.StatusOK, gin.H{
			// 		"error": "无效的id",
			// 	})
			// }

			// if err := DB.Model(&Todo{}).Where("id=?", id).Update("status", 1).Error; err != nil {
			// 	ctx.JSON(http.StatusOK, gin.H{
			// 		"error": "修改失败",
			// 	})

			// }
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})

			}

			fmt.Println()
			fmt.Println("===========")
			fmt.Println(ctx)
			fmt.Println("===========")
			fmt.Println()

			fmt.Println("bind前todo=", todo)
			ctx.BindJSON(&todo)
			fmt.Println("bind后todo=", todo)

			if err = DB.Save(&todo).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, todo)
			}
		})
		v1Group.DELETE("/todo/:id", func(ctx *gin.Context) {
			id, ok := ctx.Params.Get("id")
			if !ok {
				ctx.JSON(http.StatusOK, gin.H{
					"error": "无效的id",
				})
			}

			if err := DB.Model(&Todo{}).Delete("id=?", id).Error; err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"error": "删除失败",
				})
			}
		})

		router.Run("127.0.0.1:8080")
	}
}
