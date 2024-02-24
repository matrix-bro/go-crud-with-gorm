package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"example/go-crud/serializers"

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
