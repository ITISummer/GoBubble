package controller

import (
	"bubblt/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 首页
func IndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// 添加任务
func AddTodo(c *gin.Context) {
	var todo models.Todo
	// 1. 获取请求中的数据
	err := c.BindJSON(&todo)
	if err != nil {
		fmt.Println("获取请求参数出现错误！")
		return
	}
	// 2. 存入数据库
	err = models.AddTodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 删除任务
func DelTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的 id"})
		return
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "成功删除！"})
	}
}

// 修改任务
func UpdateTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的 id"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// 得到所有任务
func GetAllTodo(c *gin.Context) {
	if todoList, err := models.GetAllTodo(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// 通过 id 得到任务
func GetTodoById(c *gin.Context) {

}
