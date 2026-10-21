package router

import (
	"gin/bubble/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	//找静态文件
	router.Static("bubble/static", "static") //第一个参数：url上的路径，就是对外暴露的路径。第二个参数是电脑上的实际路径
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.IndexHandler)
	v1Group := router.Group("/v1")
	{
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有
		v1Group.GET("/todo", controller.GetTodoList)
		//查看单个
		v1Group.GET("/todo/:id", func(ctx *gin.Context) {

		})

		//修改
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)

	}
	return router

}
