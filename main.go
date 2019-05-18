package main

import (
	"github.com/chanhung34/todo_list/helper"
	"github.com/chanhung34/todo_list/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/testing", userRegister)

	err := r.Run()
	if err != nil {
		log.Println(err.Error())
	}
}

func userRegister(c *gin.Context) {
	var urr model.UserRegisterRequest
	if c.ShouldBind(urr) != nil {
		c.String(http.StatusOK, helper.INVALID_INPUT_REQUEST_MSG)
	}


}
