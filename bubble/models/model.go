package models

import "gin/bubble/database"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateATodo(todo *Todo) (err error) {
	err = database.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todoList []*Todo, err error) {
	if err = database.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return todoList, nil
}

func GetATodo(id string) (todo *Todo, err error) {
	if err := database.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func UpdateATodo(todo *Todo) (err error) {
	if err := database.DB.Save(&todo).Error; err != nil {
		return err
	}
	return nil
}

func DeleteATodo(id string) (err error) {
	err = database.DB.Model(&Todo{}).Delete("id=?", id).Error
	return
}
