package repository

import "todo_list/model"

type TodoItems interface {
	GetUserTodoList(userId int, page int, limit int) ([]*model.TodoItem, error)
	CreateTodoItem(userId int, title string) (*model.TodoItem, error)
	UpdateTodoItem(userId int, itemID int, title string, status string) (*model.TodoItem, error)
	DeleteTodoItem(userId int, itemID int, title string, status string) (*model.TodoItem, error)
}
