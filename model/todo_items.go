package model

import "time"

const done = "done"
const todo = "todo"
const doing = "doing"

type TodoItems struct {
	ID        int        `json:""  gorm:"column:ID; primary_key; AUTO_INCREMENT"`
	UserID    int        `json:"user_id" gorm:"column:UserID"`
	CreatedAt *time.Time `json:"create_at" gorm:"column:CreatedAt"`
	UpdateAt  *time.Time `json:"update_at" gorm:"column:UpdateAt"`
	Title     string     `json:"title" gorm:"column:Title; type:nvarchar(100); not null"`
	Status    string     `json:"status" gorm:"column:Status; type:nvarchar(10); not null"`
}
