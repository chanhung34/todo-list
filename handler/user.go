package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"net/http"
	"todo_list/model"
)

type User interface {
	Register(c *gin.Context)
	Auth(c *gin.Context)
}

type commonUser struct {
	logger *logrus.Logger
	ctx    *context.Context
	gormDB *gorm.DB
}

func NewUser(logger *logrus.Logger, ctx context.Context, db *gorm.DB) *commonUser {
	return &commonUser{
		logger: logger,
		ctx:    &ctx,
		gormDB: db,
	}
}

func (handler *commonUser) Register(c *gin.Context) {
	var urr model.UserRegisterRequest
	if err := c.ShouldBindJSON(&urr); err != nil {
		c.JSON(http.StatusBadRequest, model.UserRegisterResponse{IsError: true, ErrorMessage: err.Error()})
		return
	}
	if urr.UserName != "hungvtc" || urr.Password != "123abc" {
		c.JSON(http.StatusOK, model.UserRegisterResponse{IsError: true, ErrorMessage: "fail to authorized"})
		return
	}
	//logger := log.N
	//l\
	c.JSON(http.StatusOK, model.UserRegisterResponse{Data: model.UserRegisterResponseData{"ahihi", "passwordd", "access token"}})
}
