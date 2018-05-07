package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/users/signup", controller.HandleSignup)
	router.POST("/users/login", controller.HandleLogin)

	router.POST("/articles", controller.HandlePostArticle)
	router.PUT("/articles/:id", controller.HandlePutArticle)
	router.DELETE("/articles/:id", controller.HandleDeleteArticle)
	router.GET("/articles", controller.HandleGetArticles)
	router.GET("/articles/:id", controller.HandleGetArticle)

	router.Run(":9345")
}
