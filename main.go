package main

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/zignd/go-blog-api/handlers"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func main() {
	router := gin.Default()

	router.GET("/posts/:title-url", handlers.GetPostsTitleUrl)
	router.GET("/posts", handlers.GetPosts)
	router.POST("/posts", handlers.PostPosts)

	router.Run(":8080")
}
