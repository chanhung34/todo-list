package model

import "time"

const STATUS_DONE = "done"
const STATUS_TODO = "todo"
const STATUS_DOING = "doing"

type TodoItem struct {
	ID        int        `json:"id"  gorm:"column:id; primary_key; AUTO_INCREMENT"`
	UserID    int        `json:"user_id" gorm:"column:user_id"`
	CreatedAt *time.Time `json:"create_at" gorm:"column:created_at"`
	UpdateAt  *time.Time `json:"update_at" gorm:"column:update_at"`
	Title     string     `json:"title" gorm:"column:title; type:nvarchar(100); not null"`
	Status    string     `json:"status" gorm:"column:status; type:nvarchar(10); not null"`
}

func (*TodoItem) GetTableName() string {
	return "todo_items"
}
