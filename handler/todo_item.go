package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
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
	accountRepository := repository.NewAccountRepository(accountStorage, handler.logger)
	acc, err := accountRepository.VerifyAuthenticate(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.AddTodoItemResponse{IsError: true, Message: err.Error()})
	}
	fmt.Printf("%v", acc)
}

func (handler *todoItemHandler) Update(c *gin.Context) {
	panic("implement me")
}

func (handler *todoItemHandler) Delete(c *gin.Context) {
	panic("implement me")
}

func (handler *todoItemHandler) Listing(c *gin.Context) {
	panic("implement me")
}

func NewTodoItemHandler(logger logrus.Logger, ctx context.Context, db *gorm.DB) *todoItemHandler {
	return &todoItemHandler{
		logger: logger,
		ctx:    &ctx,
		gormDB: db,
	}
}
