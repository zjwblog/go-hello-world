package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/test"
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open [%s] failed, err: %v\n", dsn, err)
		return
	}
	err = db.Ping() // 尝试链接数据库
	if err != nil {
		fmt.Printf("Ping [%s] failed, err: %v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功!")
}
