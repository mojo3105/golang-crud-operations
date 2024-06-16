package controllers

import (
	"fmt"
	"go-crud/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// GET /books
// Get all books
func FindBooks(c *gin.Context) {
	id := c.Param("id") 
	if id != "" {
		var books models.Book
		models.DB.First(&books, id)
		c.JSON(http.StatusOK, gin.H{"data": books})
	} else {
		var books []models.Book
		models.DB.Find(&books)
		c.JSON(http.StatusOK, gin.H{"data": books})
	}
	
}


// POST /books
// Create new book
func CreateBook(c *gin.Context) {
	// Validate input
	//get response data
	var input models.Book
	//return if got error while binding response
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// validate.RegisterValidation("custValidation", CustomValidation)
	// Validate the input data
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		validationError := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			// add the validation error messages to the map
			name := err.Field()
			fmt.Println(name)
			tag := err.Tag()
			fmt.Println(tag)
			msg := strings.Split(err.Error(), ": ")[1]
			fmt.Println(msg)
			validationError[name] = msg
		}
		c.JSON(http.StatusBadRequest, gin.H{"Error": validationError})
		return
	}

	// Create book
	// book := models.Book{Title: input.Title, Author: input.Author}
	result := models.DB.Create(&input)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Statement.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": input})
  }


// PUT /book
//update book
func UpdateBook(c *gin.Context) {
	//get id from url
	id := c.Param("id")
	
	//get response data
	var books models.Book
	//return if got error while binding response
	if err := c.ShouldBindJSON(&books); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// validate.RegisterValidation("custValidation", CustomValidation)
	// Validate the input data
	validate := validator.New()
	if err := validate.Struct(books); err != nil {
		validationError := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			// add the validation error messages to the map
			name := err.Field()
			fmt.Println(name)
			tag := err.Tag()
			fmt.Println(tag)
			msg := strings.Split(err.Error(), ": ")[1]
			fmt.Println(msg)
			validationError[name] = msg
		}
		c.JSON(http.StatusBadRequest, gin.H{"Error": validationError})
		return
	}

	//find the book
	result := models.DB.First(&books, id)
	//return if got error while finding id
	if result.Error != nil{
		fmt.Println(id) 
		response := fmt.Sprintf("Given id:%v not found!", id)
		c.JSON(http.StatusNotFound, gin.H{"Error": response})
		return
	}

	//update the data 
	result_update := models.DB.Model(&books).Updates(models.Book{Title: books.Title, Author: books.Author})

	//return if error occured while updating
	if result_update.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result_update.Statement.Error})
		return
	}

	//finally return the success response
	c.JSON(http.StatusOK, gin.H{"Success": books})
}


//DELETE /books
func DeleteBooks(c *gin.Context) {
	//get id from url
	id := c.Param("id")

	//delete the book
	result := models.DB.Delete(&models.Book{}, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": fmt.Sprintf("Provided book-id %v not found!", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success": fmt.Sprintf("Provided book with id %v deleted!", id)})
}