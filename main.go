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
	//initial
	r := gin.Default()
	api := r.Group("/api/v1")
	gorm, err := GORMInit()
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	userHandler := handler.NewUser(*logger, context.Background(), gorm)
	todoItemHandler := handler.NewTodoItemHandler(*logger, context.Background(), gorm)
	// no authentication endpoints
	{
		api.POST("/register", userHandler.Register)
		api.POST("/auth", userHandler.Auth)
	}
	// routing
	// basic authentication endpoints
	{
		basicAuth := api.Group("/")
		//basicAuth.Use(AuthenticationRequired())
		//{
		basicAuth.POST("/todo_items/create", todoItemHandler.Create)
		basicAuth.PUT("/todo_items/update", userHandler.Register)
		basicAuth.GET("/todo_items/list/:user_id/:page", userHandler.Register)
		basicAuth.DELETE("/todo_items/delete/:user_id", userHandler.Register)
		//}
	}

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

//func AuthenticationRequired(auths ...string) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		session := http.Header{}.Default(c)
//		user := session.Get("user")
//		if user == nil {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
//			c.Abort()
//			return
//		}
//		if len(auths) != 0 {
//			authType := session.Get("authType")
//			if authType == nil || !funk.ContainsString(auths, authType.(string)) {
//				c.JSON(http.StatusForbidden, gin.H{"error": "invalid request, restricted endpoint"})
//				c.Abort()
//				return
//			}
//		}
//		// add session verification here, like checking if the user and authType
//		// combination actually exists if necessary. Try adding caching this (redis)
//		// since this middleware might be called a lot
//		c.Next()
//	}
//}
