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

func GetAuthorByID(id string) (*models.Author, error) {
	authorId, err := utils.ConvertStringToUint(id)

	if err != nil {
		return nil, err
	}
	var author models.Author
	checkAuthor := initializers.DB.First(&author, authorId).Error

	if checkAuthor != nil {
		return nil, checkAuthor
	}

	return &author, nil
}

func CreateBook(c *gin.Context) {
	author, err := GetAuthorByID(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var bookSerializer serializers.BookSerializer

	if err := c.ShouldBindJSON(&bookSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Data"})
		return
	}

	newBook := models.Book{Title: bookSerializer.Title, Description: bookSerializer.Description, AuthorID: author.ID}
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
