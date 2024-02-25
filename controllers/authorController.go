package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"example/go-crud/serializers"
	"example/go-crud/utils"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	var authorSerializer serializers.AuthorSerializer

	if err := c.ShouldBindJSON(&authorSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Data"})
		return
	}

	author := models.Author{FirstName: authorSerializer.FirstName, LastName: authorSerializer.LastName}

	result := initializers.DB.Create(&author)

	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error creating new author."})
		return
	}

	c.IndentedJSON(201, gin.H{"data": authorSerializer, "message": "Created new author successfully."})
}

func CreateBook(c *gin.Context) {
	authorId, err := utils.ConvertStringToUint(c.Param("id"))

	if err != "" {
		c.IndentedJSON(200, gin.H{"error": err})
		return
	}

	checkAuthor := initializers.DB.First(&models.Author{}, authorId)

	if checkAuthor.Error != nil {
		c.IndentedJSON(200, gin.H{"data": "Author not found"})
		return
	}

	var bookSerializer serializers.BookSerializer

	if err := c.ShouldBindJSON(&bookSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Data"})
		return
	}

	newBook := models.Book{Title: bookSerializer.Title, Description: bookSerializer.Description, AuthorID: authorId}
	result := initializers.DB.Create(&newBook)

	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error creating new book."})
		return
	}

	c.IndentedJSON(201, gin.H{"data": bookSerializer, "message": "Created new book successfully."})

}

func GetAllAuthors(c *gin.Context) {
	var authors []serializers.AuthorSerializer
	result := initializers.DB.Model(&models.Author{}).Find(&authors)
	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error retrieving authors."})
		return
	}
	c.IndentedJSON(201, gin.H{"data": authors, "message": "Author retrieved successfully."})
}

func GetAllBooks(c *gin.Context) {
	var books []serializers.BookSerializer
	result := initializers.DB.Model(&models.Book{}).Find(&books)
	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error retrieving books."})
		return
	}
	c.IndentedJSON(201, gin.H{"data": books, "message": "Books retrieved successfully."})
}
