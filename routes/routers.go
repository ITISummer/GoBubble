package routes

import (
	"bubblt/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	r := gin.Default()
	// 告诉 gin 去指定目录下寻找静态文件
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", controller.IndexController)

	// api v1 路由组
	v1Group := r.Group("v1")

	// 添加一个代办
	v1Group.POST("/todo", controller.AddTodo)
	// 删除某一个代办
	v1Group.DELETE("/todo/:id", controller.DelTodoById)
	// 修改某一个代办
	v1Group.PUT("/todo/:id", controller.UpdateTodo)
	// 查询所有代办
	v1Group.GET("/todo", controller.GetAllTodo)
	// 查询某一个代办
	v1Group.GET("/todo/:id", controller.GetTodoById)
	return r
}
