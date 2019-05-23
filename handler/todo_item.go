package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"todo_list/model"
	"todo_list/repository"
	"todo_list/storage"
)

type TodoItemHandler interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Listing(c *gin.Context)
}
type todoItemHandler struct {
	logger logrus.Logger
	ctx    *context.Context
	gormDB *gorm.DB
}

func (handler *todoItemHandler) Create(c *gin.Context) {
	var uar model.AddTodoItemRequest
	// Get the JSON body and decode into credentials
	if err := c.ShouldBindJSON(&uar); err != nil {
		c.JSON(http.StatusBadRequest, model.AddTodoItemResponse{IsError: true, Message: err.Error()})
		return
	}
	accountStorage := storage.NewCustomerStorage(handler.gormDB, handler.logger)
	todoItemStorage := storage.NewTodoItem(handler.gormDB, handler.logger)
	accountRepository := repository.NewAccountRepository(accountStorage, handler.logger)
	todoItemRepository := repository.NewTodoItemRepository(todoItemStorage, handler.logger)
	acc, err := accountRepository.VerifyAuthenticate(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.AddTodoItemResponse{IsError: true, Message: err.Error()})
		return
	}
	if acc == nil {
		c.JSON(http.StatusUnauthorized, model.AddTodoItemResponse{IsError: true, Message: "Loggin account is invalid"})
		return
	}
	todoItem, err2 := todoItemRepository.CreateTodoItem(acc.ID, uar.Tittle)
	if err2 != nil {
		c.JSON(http.StatusOK, model.AddTodoItemResponse{IsError: true, Message: err2.Error()})
		return
	}
	c.JSON(http.StatusOK, model.AddTodoItemResponse{TodoItem: todoItem})
	return

}

func (handler *todoItemHandler) Update(c *gin.Context) {
	panic("implement me")
}

func (handler *todoItemHandler) Delete(c *gin.Context) {
	panic("implement me")
}

func (handler *todoItemHandler) Listing(c *gin.Context) {
	// Get the JSON body and decode into credentials
	pageNumb, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, model.GetTodoItemsResponse{IsError: true, Message: err.Error()})
		return
	}
	accountStorage := storage.NewCustomerStorage(handler.gormDB, handler.logger)
	todoItemStorage := storage.NewTodoItem(handler.gormDB, handler.logger)
	accountRepository := repository.NewAccountRepository(accountStorage, handler.logger)
	todoItemRepository := repository.NewTodoItemRepository(todoItemStorage, handler.logger)
	acc, err := accountRepository.VerifyAuthenticate(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.GetTodoItemsResponse{IsError: true, Message: err.Error()})
		return
	}
	if acc == nil {
		c.JSON(http.StatusUnauthorized, model.GetTodoItemsResponse{IsError: true, Message: "Loggin account is invalid"})
		return
	}
	todoItems, err2 := todoItemRepository.GetUserTodoList(acc.ID, pageNumb, 5)
	if err2 != nil {
		c.JSON(http.StatusOK, model.GetTodoItemsResponse{IsError: true, Message: err2.Error()})
		return
	}
	c.JSON(http.StatusOK, model.GetTodoItemsResponse{TodoItems: todoItems})
	return
}

func NewTodoItemHandler(logger logrus.Logger, ctx context.Context, db *gorm.DB) *todoItemHandler {
	return &todoItemHandler{
		logger: logger,
		ctx:    &ctx,
		gormDB: db,
	}
}
