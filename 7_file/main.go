package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/index", func(ctx *gin.Context) {

		//单个文件
		file, err := ctx.FormFile("file")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)

		ext := path.Ext(file.Filename) //获取后缀
		newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
		filepath := fmt.Sprintf("./%s", newFileName)
		ctx.SaveUploadedFile(file, filepath)
		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s'uploaded!", file.Filename),
		})
	})
	router.Run("127.0.0.1:8080")
}
