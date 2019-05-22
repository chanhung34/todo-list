package repository

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
	"todo_list/model"
	"todo_list/storage"
)

type AccountRepository interface {
	Validate() (map[string]interface{}, bool)
	SignUp(ctx context.Context, newAccount *model.Account) (*model.Account, error)
	GetUser(ctx context.Context, userId uint) *model.Account
	JwtAuthentication(ctx context.Context, next http.Handler) http.Handler
}

type account struct {
	storage storage.Account
	logger  logrus.Logger
}

func NewAccountRepository(sa storage.Account, logger logrus.Logger) *account {
	return &account{
		storage: sa,
		logger:  logger,
	}
}

func (repo account) Validate() (map[string]interface{}, bool) {
	panic("implement me")
}

func (repo account) SignUp(ctx context.Context, newAccount *model.Account) (*model.Account, error) {
	isExistAccount, err := repo.storage.ExistUserName(ctx, newAccount.UserName)
	if err != nil {
		return nil, err
	}
	if isExistAccount {
		return nil, errors.New("User already ")
	}
	now := time.Now()
	newAccount.CreatedAt = &now
	newAccount.UpdatedAt = &now
	newAccount, err2 := repo.storage.Create(ctx, newAccount)
	if err2 != nil {
		return nil, err2
	}
	return newAccount, nil

}

func (repo account) GetUser(ctx context.Context, userId uint) *model.Account {
	panic("implement me")
}

func (repo account) JwtAuthentication(ctx context.Context, next http.Handler) http.Handler {
	panic("implement me")
}
