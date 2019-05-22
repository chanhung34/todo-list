package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}
type Account struct {
	ID        int        `json:"user_id"  gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserName  string     `json:"user_name" gorm:"column:user_name; type:nvarchar(100); not null"`
	Password  string     `json:"password" gorm:"column:password;  type:nvarchar(100); not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (*Account) TableName() string {
	return "user_accounts"
}
