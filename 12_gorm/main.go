package main

import (
	"fmt"
	"gin/12_gorm/global"
	"log"
	"time"

	"gorm.io/gorm"
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
	DeletedAt gorm.DeletedAt //有这个字段则可以软删除,Delete为软删除,unscoped().Delete为物理删除
}

func migrate() {
	err := global.DB.AutoMigrate(&UserModel{})
	if err != nil {
		fmt.Println("表迁移失败：", err)
		return
	}
	fmt.Println("表迁移成功")
}

func insert() {
	// migrate()

	//回填式创建
	user := UserModel{
		Name: "张三1",
		Age:  25,
	}
	// global.DB.Create(&user)
	// fmt.Println(user.ID, user.Age, user.Name, user.CreatedAt)
	// err := global.DB.Create(&UserModel{
	// 	Name: "枫枫",
	// 	Age:  25,
	// }).Error
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var userList = []UserModel{
	// 	{Name: "王五"},
	// 	{Name: "李四"},
	// }

	err := global.DB.Create(&user).Error
	fmt.Println("创建完成。。。")
	if err != nil {
		fmt.Println(err)
		return
	}

}
func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	fmt.Println("运行create钩子。。。")
	return nil
}

func query() {
	// var userList []UserModel
	// global.DB.Find(&userList, "age>=?", 25)
	// fmt.Println(userList)

	var user UserModel
	// global.DB.Take(&user, "id=?", 3)
	// fmt.Println(user)

	global.DB.Debug().First(&user, 5) //默认按照主键查询

	fmt.Println(user)
	global.DB.Last(&user, 6) //查完5后，user里面id=5，再次查询时会把id=5放入查询条件里边，导致查询id=5&id=6从而error
	fmt.Println(user)
}

func updateSave() {
	var user UserModel
	user.ID = 3
	user.Name = "张四"
	user.CreatedAt = time.Now()
	global.DB.Save(&user)
	fmt.Println(user)
}
func update() {
	var user = UserModel{ID: 3}
	// user.ID = 1
	// global.DB.Take(&test, 1)
	// fmt.Println(test.Age)
	global.DB.Model(&user).Update("age", 3)
	fmt.Println(user.Age)
}

func delete() {
	// var user UserModel
	// user.ID = 8
	// global.DB.Delete(&user)

	// global.DB.Delete(&UserModel{}, 7)
	global.DB.Delete(&UserModel{}, "id=?", 9)
}

// 迁移--把结构体转化为数据库
func Migrate() {
	err := global.DB.AutoMigrate(&UserModel{})
	if err != nil {
		log.Fatalf("迁移失败：%s", err)
	}
	log.Fatalf("迁移成功")
}

func scan() {
	// var nameList []string

	// global.DB.Model(UserModel{}).Select("name").Scan(&nameList)
	// fmt.Println(nameList)
	type U struct {
		Name string
		Age  int
	}
	var u []U
	global.DB.Debug().Model(&UserModel{}).Scan(&u)
	fmt.Println(u)
}

// 分页查询
func offset() {
	var UserList []UserModel
	global.DB.Limit(2).Offset(0).Find(&UserList)
	fmt.Println(UserList)
	global.DB.Limit(2).Offset(2).Find(&UserList)
	fmt.Println(UserList)
}

// 分页查询——按照page
func offset_page() {

	limit := 2
	var UserList []UserModel
	// //第一页
	page := 1
	global.DB.Limit(limit).Offset((page - 1) * limit).Find(&UserList)
	fmt.Println(UserList)
	fmt.Println("===================")
	// //第二页
	// page = 2
	// global.DB.Limit(limit).Offset((page - 1) * limit).Find(&UserList)
	// fmt.Println(UserList)
	page++
	for len(UserList) != 0 {
		global.DB.Limit(limit).Offset((page - 1) * limit).Find(&UserList)
		fmt.Println(UserList)
		fmt.Println("===================")
		page++
	}
}
func main() {

	global.Connect()
	// Migrate()
	// update()

	// delete()
	// var user UserModel
	// global.DB.Unscoped().Take(&user, "name=?", "aa")
	// fmt.Println(user)

	// scan()
	// offset()
	offset_page()
	// var users []UserModel
	// for i := 1; i <= 10; i++ {
	// 	users = append(users, UserModel{
	// 		Name: fmt.Sprintf("测试用户_%02d", i), // 生成如 "测试用户_01"，保证 unique 且不超 16 字符
	// 		Age:  18 + i,                      // 年龄 19 到 28
	// 		// CreatedAt 不需要填，GORM 会自动填入当前时间
	// 		// DeletedAt 不需要填，默认为 NULL（未删除）
	// 	})
	// }
	// global.DB.Create(&users)
}
