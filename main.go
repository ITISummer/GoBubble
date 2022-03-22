package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

// 定义全局变量
var (
	DB *gorm.DB
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func main() {
	// 连接数据库
	err := initMysql()
	if err != nil {
		panic(err)
	}
	// 若程序退出则关闭数据库连接
	defer DB.Close()
	// 模型绑定 - 会创建 todos 数据库表
	DB.Debug().AutoMigrate(&Todo{})

	r := gin.Default()
	// 告诉 gin 去指定目录下寻找静态文件
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// api v1 路由组
	v1Group := r.Group("v1")

	// 添加一个代办
	v1Group.POST("/todo", func(c *gin.Context) {
		var todo Todo
		// 1. 获取请求中的数据
		err := c.BindJSON(&todo)
		if err != nil {
			fmt.Println("获取请求参数出现错误！")
			return
		}
		// 2. 存入数据库
		if err = DB.Create(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	})
	// 删除某一个代办
	v1Group.DELETE("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": "无效的 id"})
			return
		}
		if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{id: "成功删除！"})
		}
	})
	// 修改某一个代办
	v1Group.PUT("/todo/:id", func(c *gin.Context) {
		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"error": "无效的 id"})
			return
		}
		var todo Todo
		if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		c.BindJSON(&todo)
		if err = DB.Save(&todo).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todo)
		}
	})
	// 查询所有代办
	v1Group.GET("/todo", func(c *gin.Context) {
		var todoList []Todo
		if err := DB.Find(&todoList).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, todoList)
		}
	})
	// 查询某一个代办
	v1Group.GET("/todo/:id", func(c *gin.Context) {

	})

	// 默认 8080 端口 - 放在路由组后
	r.Run()
}

// 初始化mysql链接
func initMysql() (err error) {
	sqlStr := "root:root@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", sqlStr)

	if err != nil {
		fmt.Println("产生错误")
		return
	}
	return DB.DB().Ping()
}
