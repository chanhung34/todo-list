package storage

type TodoItemStorage interface {
	UpdateTodoItem(userId int, item *model.TodoItem) (*model.TodoItem, error)
	CreateTodoItem(item *model.TodoItem) (*model.TodoItem, error)
	GetTodoItemById(itemId int) (*model.TodoItem, error)
	GetTodoItemsByUserId(itemId int, page int, limit int) ([]*model.TodoItem, error)
	DeleteTodoItem(itemId int) error
}
type todoItem struct {
}

func (storage *todoItem) UpdateTodoItem(userId int, item *model.TodoItem) (*model.TodoItem, error) {

}
func (storage *todoItem) CreateTodoItem(item *model.TodoItem) (*model.TodoItem, error) {}
func (storage *todoItem) GetTodoItemById(itemId int) (*model.TodoItem, error)          {}
func (storage *todoItem) GetTodoItemsByUserId(itemId int, page int, limit int) ([]*model.TodoItem, error) {
}
func (storage *todoItem) DeleteTodoItem(itemId int) error {}
