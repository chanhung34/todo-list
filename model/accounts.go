package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
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
	gorm.Model
	UserID    int        `json:"user_id"  gorm:"column:UserID; primary_key; AUTO_INCREMENT"`
	UserName  string     `json:"user_name" gorm:"column:UserID; type:nvarchar(100); not null"`
	Password  string     `json:"password" gorm:"column:Password;  type:nvarchar(100); not null"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:CreatedAt"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:UpdatedAt"`
	Token     string     `json:"token";sql:"-"`
}
