package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func main() {
	// 创建数据库 TODO

	r := gin.Default()
	// 告诉 gin 去指定目录下寻找静态文件
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	// 默认 8080 端口
	r.Run()

	// api v1
	v1Group := r.Group("v1")
	{
		// 添加一个代办
		v1Group.POST("/todo", func(c *gin.Context) {

		})
		// 删除某一个代办
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
		// 修改某一个代办
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})
		// 查询所有代办
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		// 查询某一个代办
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
	}

}
