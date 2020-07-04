package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ravindra031/golang_gin_rest/controllers"
	"github.com/ravindra031/golang_gin_rest/models"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"data": "hello world"})
	})

	models.ConnectDataBase()

	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
