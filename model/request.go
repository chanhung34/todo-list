package model

//UserRegister
type UserRegisterRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterResponse struct {
	IsError bool                     `json:"is_error"`
	Message string                   `json:"message"`
	Data    UserRegisterResponseData `json:"data"`
}
type UserRegisterResponseData struct {
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

// user auth
type UserAuthRequest struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type UserAuthResponse struct {
	IsError bool     `json:"is_error"`
	Message string   `json:"message"`
	Data    UserAuth `json:"data"`
}
type UserAuth struct {
	UserName    string `json:"user_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	AccessToken string `json:"access_token"`
}

// Add todo-item
type AddTodoItemRequest struct {
	UserId int    `json:"user_id" binding:"required"`
	Tittle string `json:"tittle" binding:"required"`
}
type AddTodoItemResponse struct {
	IsError  bool      `json:"is_error"`
	Message  string    `json:"message"`
	TodoItem *TodoItem `json:"todo_item"`
}

// Add todo-item
type UpdateTodoItemRequest struct {
	ItemId int    `json:"item_id" binding:"required"`
	Tittle string `json:"tittle"`
	Status string `json:"status"`
}
type UpdateTodoItemResponse struct {
	IsError  bool      `json:"is_error"`
	Message  string    `json:"message"`
	TodoItem *TodoItem `json:"todo_item"`
}

// delete todo-item
type DeleteTodoItemRequest struct {
	UserId int `json:"user_id" binding:"required"`
}
type DeleteTodoItemResponse struct {
	IsError bool   `json:"is_error"`
	Message string `json:"message"`
}

// Get todo-items
type GetTodoItemsRequest struct {
	UserId int `json:"user_id"`
	Page   int `json:"page"`
}
type GetTodoItemsResponse struct {
	IsError   bool        `json:"is_error"`
	Message   string      `json:"message"`
	TodoItems []*TodoItem `json:"todo_items"`
}
