package main

import (
	"bubblt/models"
	"bubblt/routes"
	"bubblt/utils"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 连接数据库
	err := utils.InitMysql()
	if err != nil {
		panic(err)
	}
	// 若程序退出则关闭数据库连接
	defer utils.CloseDB()
	// 模型绑定 - 会创建 todos 数据库表
	utils.DB.Debug().AutoMigrate(&models.Todo{})
	r := routes.SetupRouters()
	// 默认 8080 端口 - 放在路由组后
	r.Run()
}
