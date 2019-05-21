package model

const done = "done"
const todo = "todo"
const doing = "doing"

type TodoItems struct {
	UserID     uint64 `json:"user_id"`
	CreateTime uint64 `json:"create_time"`
	UpdateTime uint64 `json:"update_time"`
	Title      uint64 `json:"title"`
	Status     string `json:"status"`
}
