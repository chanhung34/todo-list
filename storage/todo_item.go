package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"todo_list/model"
)

type TodoItemStorage interface {
	UpdateTodoItem(item *model.TodoItem) (*model.TodoItem, error)
	CreateTodoItem(item *model.TodoItem) (*model.TodoItem, error)
	GetTodoItemById(itemId int) (*model.TodoItem, error)
	GetTodoItemsByUserId(userId int, page int, limit int) ([]*model.TodoItem, error)
	DeleteTodoItem(todoItem *model.TodoItem) (bool, error)
}
type todoItem struct {
	db     *gorm.DB
	logger logrus.Logger
}

func NewTodoItem(db *gorm.DB, logger logrus.Logger) *todoItem {
	return &todoItem{
		db:     db,
		logger: logger,
	}
}
func (storage *todoItem) UpdateTodoItem(item *model.TodoItem) (*model.TodoItem, error) {
	err := storage.db.Model(&item).UpdateColumns(item).Error
	//return false, nil
	return item, err
}
func (storage *todoItem) CreateTodoItem(item *model.TodoItem) (*model.TodoItem, error) {
	err := storage.db.New().Omit("VersionNo").Create(item).Error
	//return false, nil
	return item, err
}
func (storage *todoItem) GetTodoItemById(itemId int) (*model.TodoItem, error) {
	var item model.TodoItem
	err := storage.db.Model(model.TodoItem{}).Where("id = ?", itemId).First(&item).Error
	//return false, nil
	return &item, err
}
func (storage *todoItem) GetTodoItemsByUserId(userId int, page int, limit int) ([]*model.TodoItem, error) {
	var items []*model.TodoItem
	err := storage.db.Model(model.TodoItem{}).Limit(limit).Offset((page-1)*limit).Where("user_id = ?", userId).Find(&items).Error
	//return false, nil
	return items, err
}
func (storage *todoItem) DeleteTodoItem(todoItem *model.TodoItem) (bool, error) {
	err := storage.db.Delete(&todoItem).Error
	if err != nil {
		return false, err
	}
	return true, err
}
