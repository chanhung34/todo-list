package repository

import (
	"github.com/sirupsen/logrus"
	"time"
	"todo_list/model"
	"todo_list/storage"
)

type TodoItems interface {
	GetUserTodoList(userId int, page int, limit int) ([]*model.TodoItem, error)
	CreateTodoItem(userId int, title string) (*model.TodoItem, error)
	UpdateTodoItem(userId int, itemID int, title string, status string) (*model.TodoItem, error)
	DeleteTodoItem(userId int, itemID int, title string, status string) (*model.TodoItem, error)
}

type todoItem struct {
	storage storage.TodoItemStorage
	logger  logrus.Logger
}

func NewTodoItemRepository(storage storage.TodoItemStorage, logger logrus.Logger) *todoItem {
	return &todoItem{
		storage: storage,
		logger:  logger,
	}
}

func (repo todoItem) GetUserTodoList(userId int, page int, limit int) ([]*model.TodoItem, error) {
	rs, err := repo.storage.GetTodoItemsByUserId(userId, page, limit)
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (repo todoItem) CreateTodoItem(userId int, title string) (*model.TodoItem, error) {
	now := time.Now()
	newTDItem := model.TodoItem{UserID: userId, Title: title, Status: model.STATUS_TODO, CreatedAt: &now, UpdateAt: &now}
	rs, err := repo.storage.CreateTodoItem(&newTDItem)
	if err != nil {
		return nil, err
	}
	return rs, nil

}

func (repo *todoItem) UpdateTodoItem(userId int, itemID int, title string, status string) (*model.TodoItem, error) {
	panic("implement me")
}

func (repo *todoItem) DeleteTodoItem(userId int, itemID int, title string, status string) (*model.TodoItem, error) {
	panic("implement me")
}
