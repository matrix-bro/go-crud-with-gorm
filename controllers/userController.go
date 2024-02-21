package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Data"})
		return
	}

	// save the user
	result := initializers.DB.Create(&newUser)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.IndentedJSON(200, gin.H{"user": newUser})
}
