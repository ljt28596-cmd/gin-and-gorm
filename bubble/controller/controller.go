package controller

import (
	"gin/bubble/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(ctx *gin.Context) {
	//从请求中把数据拿出来
	var todo models.Todo
	ctx.BindJSON(&todo)
	//存入数据库
	//返回数据
	err := models.CreateATodo(&todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}

func GetTodoList(ctx *gin.Context) {
	//查看todos表里的所有数据

	todoList, err := models.GetAllTodo()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, todoList)
	}
}

func UpdateTodo(ctx *gin.Context) {
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
		return
	}

	todo, err := models.GetATodo(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return

	}

	ctx.BindJSON(todo)
	err = models.UpdateATodo(todo)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
}
func DeleteTodo(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
		return
	}

	if err := models.DeleteATodo(id); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "删除失败",
		})
	}
}
