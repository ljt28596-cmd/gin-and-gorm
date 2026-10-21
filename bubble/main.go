package main

import (
	"gin/bubble/database"
	"gin/bubble/models"
	"gin/bubble/router"
)

// GORM 核心库

func main() {

	err := database.InitMySQL()
	if err != nil {
		panic(err)
	}

	database.DB.AutoMigrate(&models.Todo{}) //转化为数据库表时，表名为todos
	router := router.SetupRouter()
	router.Run("127.0.0.1:8080")

}
