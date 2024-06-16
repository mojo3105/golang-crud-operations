package main

import (
	"fmt"
	"go-crud/controllers"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
    if err != nil {
        panic("Error loading .env file")
    }
}

func main() {
	fmt.Println("Starting development server")
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
	})


	models.ConnectDatabase()

	r.GET("/books/:id", controllers.FindBooks)
	r.GET("/books/", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook) 
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBooks)
	r.Run()
	fmt.Println("run finished")
}