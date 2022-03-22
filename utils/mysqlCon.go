package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

// 定义全局变量
var (
	DB *gorm.DB
)

// 初始化mysql链接
func InitMysql() (err error) {
	sqlStr := "root:root@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", sqlStr)

	if err != nil {
		fmt.Println("产生错误")
		return
	}
	return DB.DB().Ping()
}

func CloseDB() {
	// 若程序退出则关闭数据库连接
	DB.Close()
}
