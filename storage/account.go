package storage

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"todo_list/model"
)

type Account interface {
	Create(ctx context.Context, account *model.Account) (*model.Account, error)
	GetByUserNameAndPassword(ctx context.Context, username string, password string) (*model.Account, error)
	GetByUserId(ctx context.Context, userId int) (*model.Account, error)
	ExistUserName(ctx context.Context, userName string) (bool, error)
}

type accountStorage struct {
	db     *gorm.DB
	logger logrus.Logger
}

func NewCustomerStorage(db *gorm.DB, logger logrus.Logger) *accountStorage {
	return &accountStorage{
		db:     db,
		logger: logger,
	}
}
func (storage accountStorage) Create(ctx context.Context, account *model.Account) (*model.Account, error) {
	err := storage.db.New().Create(&account).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (storage accountStorage) ExistUserName(ctx context.Context, userName string) (bool, error) {
	var count = 0
	err := storage.db.Model(model.Account{}).Where("user_name = ?", userName).Count(&count).Error
	//return false, nil
	return count > 0, err
}

func (storage accountStorage) GetByUserNameAndPassword(ctx context.Context, username string, password string) (*model.Account, error) {
	var user model.Account
	err := storage.db.Model(model.Account{}).Where("user_name = ? AND password = ?", username, password).First(&user).Error
	//return false, nil
	return &user, err
}

func (storage accountStorage) GetByUserId(ctx context.Context, userId int) (*model.Account, error) {
	var user *model.Account
	err := storage.db.Model(model.Account{}).Where("id = ?", userId).First(&user).Error
	//return false, nil
	return user, err
}
