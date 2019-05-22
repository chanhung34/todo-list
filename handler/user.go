package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
	"todo_list/model"
	"todo_list/repository"
	"todo_list/storage"
)

type User interface {
	Register(c *gin.Context)
	Auth(c *gin.Context)
}

type commonUser struct {
	logger logrus.Logger
	ctx    *context.Context
	gormDB *gorm.DB
}

func NewUser(logger logrus.Logger, ctx context.Context, db *gorm.DB) *commonUser {
	return &commonUser{
		logger: logger,
		ctx:    &ctx,
		gormDB: db,
	}
}

func (handler *commonUser) Register(c *gin.Context) {
	var urr model.UserRegisterRequest
	if err := c.ShouldBindJSON(&urr); err != nil {
		c.JSON(http.StatusBadRequest, model.UserRegisterResponse{IsError: true, Message: err.Error()})
		return
	}
	accountStorage := storage.NewCustomerStorage(handler.gormDB, handler.logger)
	accountRepository := repository.NewAccountRepository(accountStorage, handler.logger)
	account, err := accountRepository.SignUp(context.Background(), &model.Account{UserName: urr.UserName, Password: urr.Password})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserRegisterResponse{IsError: true, Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, model.UserRegisterResponse{Data: model.UserRegisterResponseData{UserName: account.UserName,
			Password: account.Password, AccessToken: "access token"}})
	}

}

func (handler *commonUser) Auth(c *gin.Context) {
	var uar model.UserAuthRequest
	if err := c.ShouldBindJSON(&uar); err != nil {
		c.JSON(http.StatusBadRequest, model.UserAuthResponse{IsError: true, Message: err.Error()})
		return
	}
	accountStorage := storage.NewCustomerStorage(handler.gormDB, handler.logger)
	accountRepository := repository.NewAccountRepository(accountStorage, handler.logger)
	account, err := accountRepository.Auth(context.Background(), &model.Account{UserName: urr.UserName, Password: urr.Password})
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserAuthResponse{IsError: true, Message: err.Error()})
	} else {
		c.JSON(http.StatusOK, model.UserAuthResponse{Data: model.UserAuthResponse{UserName: account.UserName,
			Password: account.Password, AccessToken: "access token"}})
	}

}
