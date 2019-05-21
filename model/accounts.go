package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}
type Account struct {
	gorm.Model
	UserID    uint64 `json:"user_id" gorm:"primary_key; AUTO_INCREMENT"`
	UserName  string `json:"user_name" gorm:"size:255"`
	Password  string `json:"password" gorm:"size:255"`
	CreatedAt uint64 `json:"created_at"`
	UpdatedAt uint64 `json:"updated_at"`
	Token     string `json:"token";sql:"-"`
}
