package todo

import (
	data "go-web/data"
	model "go-web/model"
)

func GetTodos() []model.Todo {
	var todos []model.Todo = make([]model.Todo, 0)
	data.Database.Find(&todos)
	return todos
}

func GetTodoById(id string) (model.Todo, bool) {
	var todo model.Todo
	result := data.Database.Find(&todo, id)
	if result.RowsAffected == 0 {
		return todo, false
	}
	return todo, true
}

func CreateTodoItem(todo model.Todo) (model.Todo, bool) {
	result := data.Database.Create(&todo)
	if result.Error != nil {
		return todo, false
	}
	return todo, true
}

func UpdateTodoItem(todo model.Todo) (model.Todo, bool) {
	result := data.Database.Save(&todo)
	if result.Error != nil {
		return todo, false
	}
	return todo, true
}

func DeleteTodoById(id string) bool {
	var todo model.Todo
	result := data.Database.Delete(&todo, id)
	return result.RowsAffected > 0
}
