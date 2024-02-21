package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"strconv"

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

func CreateProfile(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid User ID"})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, userID)

	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "User not found."})
		return
	}

	req := new(models.Profile)
	err = c.ShouldBindJSON(&req)

	if err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Profile Data"})
		return
	}

	user_profile := models.Profile{Phone: req.Phone, Address: req.Address, UserID: uint(userID)}

	profile := initializers.DB.Create(&user_profile)

	if profile.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error creating Profile"})
		return
	}

	c.IndentedJSON(200, gin.H{"message": user_profile})
}
