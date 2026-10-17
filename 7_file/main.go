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

	router.POST("/single_file", func(ctx *gin.Context) {

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
	router.POST("/multi_files", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		fmt.Println("==================")
		fmt.Println("form=", form)
		fmt.Println("==================")
		fmt.Println("form.File=", form.File)
		fmt.Println("==================")
		fmt.Println("form.File['file']=", form.File["file"])
		fmt.Println("==================")

		files := form.File["file"] //去请求体(body)里找一个key叫"file"的上传项，可以更改

		for index, file := range files {
			log.Printf("%d--%s\n", index, file.Filename)
			ext := path.Ext(file.Filename)
			newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
			filepath := fmt.Sprintf("./%s", newFileName)

			ctx.SaveUploadedFile(file, filepath)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	router.Run("127.0.0.1:8080")
}
