package models

import (
	"bubblt/utils"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func AddTodo(todo *Todo) (err error) {
	if err := utils.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err := utils.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = utils.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(todo *Todo) (err error) {
	if err = utils.DB.Save(&todo).Error; err != nil {
		return
	}
	return
}

func DeleteATodo(id string) (err error) {
	err = utils.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
