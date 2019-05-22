package storage

import (
	"context"
	"gitlab.sendo.vn/core/golang-sdk/new/logger"
	"model"
)

type CustomerStorage interface {
	GetAccountByUserNameAndPassword(username string, password string) (*model.Account, error)
	GetAccountByUserID(userId int) (*model.Account, error)
	CreateAccount(customer *model.Account) (*model.Account, error)
}

type customerStorage struct {
	db     *gorm.DB
	logger logger.Logger
	ctx    context.Context
}

func NewCustomerStorage(db *gorm.DB, logger logger.Logger, ctx context.Context) *customerStorage {
	return &customerStorage{
		db:     db,
		logger: logger,
		ctx:    ctx,
	}
}
