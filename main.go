package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"log"
	"todo_list/handler"
)

func main() {
	r := gin.Default()
	gorm, err := GORMInit()
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	// routing
	userHandler := handler.NewUser(*logger, context.Background(), gorm)
	r.POST("/register", userHandler.Register)
	r.POST("/auth", userHandler.Register)
	r.POST("/todo_items/create", userHandler.Register)
	r.PUT("/todo_items/update", userHandler.Register)
	r.GET("/todo_items/list/:user_id/:page", userHandler.Register)
	r.DELETE("/todo_items/delete/:user_id", userHandler.Register)

	err = r.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func GORMInit() (*gorm.DB, error) {
	//todo: Setup env config and use .env file when config empty
	db, err := gorm.Open("mysql", "user:password@tcp(127.0.0.1:3306)/db?charset=utf8&parseTime=True")
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return db, err
}
