package model

type UserRegisterRequest struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterResponse struct {
	IsError      bool                     `json:"is_error"`
	ErrorMessage string                   `json:"error_message"`
	Data         UserRegisterResponseData `json:"data"`
}
type UserRegisterResponseData struct {
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}
