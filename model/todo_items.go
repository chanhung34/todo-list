package model

import "time"

const done = "done"
const todo = "todo"
const doing = "doing"

type TodoItem struct {
	ID        int        `json:"id"  gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserID    int        `json:"user_id" gorm:"column:user_id"`
	CreatedAt *time.Time `json:"create_at" gorm:"column:created_at"`
	UpdateAt  *time.Time `json:"update_at" gorm:"column:update_at"`
	DeleteAt  *time.Time `json:"delete_at" gorm:"column:delete_at"`
	Title     string     `json:"title" gorm:"column:title; type:nvarchar(100); not null"`
	Status    string     `json:"status" gorm:"column:status; type:nvarchar(10); not null"`
}

func (*TodoItem) GetTableName() string {
	return "todo_items"
}
