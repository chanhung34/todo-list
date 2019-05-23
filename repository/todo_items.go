package repository

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
	"todo_list/model"
	"todo_list/storage"
)

type TodoItems interface {
	GetUserTodoList(userId int, page int, limit int) ([]*model.TodoItem, error)
	CreateTodoItem(userId int, title string) (*model.TodoItem, error)
	UpdateTodoItem(userId int, itemID int, title string, status string) (*model.TodoItem, error)
	DeleteTodoItem(userId int, itemID int) error
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
	todoItem, err := repo.storage.GetTodoItemById(itemID)
	if err != nil {
		return nil, err
	}
	if userId != todoItem.UserID {
		return nil, errors.New("you can't remove other's todo item")
	}
	if status != model.STATUS_TODO && status != model.STATUS_DONE && status != model.STATUS_DOING {
		return nil, errors.New(fmt.Sprintf("We not support status : %s", status))
	}
	now := time.Now()
	if title != "" {
		todoItem.Title = title
	}
	if status != "" {
		todoItem.Status = status
	}

	todoItem.UpdateAt = &now
	rs, err := repo.storage.UpdateTodoItem(todoItem)
	if err != nil {
		return nil, err
	}
	return rs, err
}

func (repo *todoItem) DeleteTodoItem(userId int, itemID int) error {
	todoItem, err := repo.storage.GetTodoItemById(itemID)
	if err != nil {
		return err
	}
	if userId != todoItem.UserID {
		return errors.New("you can't remove other's todo item")
	}
	isSuccess, err := repo.storage.DeleteTodoItem(todoItem)
	if err != nil {
		return err
	}
	if isSuccess == false {
		return errors.New(fmt.Sprintf("Failed to delete todo item: %d", itemID))
	}
	return nil
}
