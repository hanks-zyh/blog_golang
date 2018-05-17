package main

import (
	"./conf"
	"./controller"
	"./model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func init() {
	//config := conf.Product
	config := conf.AppConfig

	connect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Mysql.Username,
		config.Mysql.Password,
		config.Mysql.Address,
		config.Mysql.Port,
		config.Mysql.DbName)
	fmt.Println(connect)
	db, err := gorm.Open("mysql", connect)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{}, &model.Article{})

	model.DB = db
}

func main() {

	config := conf.AppConfig
	gin.SetMode(config.Mode)
	router := gin.Default()

	router.POST("/users/signup", controller.HandleSignup)
	router.POST("/users/login", controller.HandleLogin)

	router.POST("/articles", controller.HandlePostArticle)
	router.PUT("/articles/:id", controller.HandlePutArticle)
	router.DELETE("/articles/:id", controller.HandleDeleteArticle)
	router.GET("/articles", controller.HandleGetArticles)
	router.GET("/articles/:id", controller.HandleGetArticle)

	router.Run(config.Server.Port)
}
