package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root1234@tcp(127.0.0.1:13306)/gorm_db_new"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("dsn格式错误: %s\n", err)

	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("数据库连接失败: %s", err)
	}
	createSQL := `CREATE TABLE IF NOT EXISTS users (
    id    BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    name  VARCHAR(50)  NOT NULL,
    age   INT          NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`

	if _, err := db.Exec(createSQL); err != nil {
		log.Fatalf("建表失败: %v", err)
	}
	// res, err := db.Exec("insert into users (name,age,email) values('张三',18,'111@gmail'),('李四',20,'222@gamil'),('王五',25,'333@gmail')")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(res)

	// rows, _ := db.Query("select id,name from users")
	// for rows.Next() {
	// 	var id int
	// 	var name string
	// 	rows.Scan(&id, &name)
	// 	fmt.Println(id, name, err)

	// }
	var id int
	var name string
	//只返回第一行
	err = db.QueryRow("select id,name from gorm_db_new.users").Scan(&id, &name)
	fmt.Println(id, name, err)

}
