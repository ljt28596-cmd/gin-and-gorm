package main

import (
	"fmt"
	"gin/12_gorm/global"
	"time"
)

type User struct {
	Id    int
	Name  string
	Email string
}
type UserModel struct {
	ID        int
	Name      string `gorm:"size:16;not null;unique"`
	Age       int    `gorm:"default:18"`
	CreatedAt time.Time
}

func migrate() {
	err := global.DB.AutoMigrate(&UserModel{})
	if err != nil {
		fmt.Println("表迁移失败：", err)
		return
	}
	fmt.Println("表迁移成功")
}

func main() {

	global.Connect()
	migrate()

}
