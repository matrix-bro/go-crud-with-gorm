package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"example/go-crud/serializers"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func AllUsers(c *gin.Context) {
	var users []models.User
	var response []serializers.UserResponse

	result := initializers.DB.Find(&users).Scan(&response)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.IndentedJSON(200, gin.H{"data": response})
}

func UserProfile(c *gin.Context) {
	var userprofiles []serializers.UserProfileResponse

	result := initializers.DB.Model(&models.User{}).Select("first_name, last_name, phone, address").
		Joins("LEFT JOIN profiles on profiles.user_id=users.id").
		Scan(&userprofiles)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.IndentedJSON(200, gin.H{"data": userprofiles})
}

func GetUserDetails(c *gin.Context) {
	userId := c.Param("id")
	var userProfile serializers.UserProfileResponse

	checkUser := initializers.DB.First(&models.User{}, userId)

	if checkUser.Error != nil {
		c.IndentedJSON(200, gin.H{"data": "User not found"})
		return
	}

	initializers.DB.Model(&models.User{}).
		Joins("LEFT JOIN profiles on users.id = profiles.user_id").
		Select("first_name, last_name, phone, address").
		Where("users.id = ?", userId).Scan(&userProfile)

	c.IndentedJSON(200, gin.H{"data": userProfile})
}

func UpdateUserDetails(c *gin.Context) {
	var updateDetails serializers.UpdateUserDetailsSerializer

	if err := c.ShouldBindJSON(&updateDetails); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Request Data"})
		return
	}

	var user models.User
	err := CheckByID(c.Param("id"), &user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	// Getting user with profile
	err = initializers.DB.Preload("Profile").First(&user, c.Param("id")).Error
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	user.FirstName = updateDetails.FirstName
	user.LastName = updateDetails.LastName

	// Check if the profile exists, create a new one if it doesn't
	if user.Profile == nil {
		user.Profile = &models.Profile{
			Phone:   updateDetails.Phone,
			Address: updateDetails.Address,
		}
	} else {
		user.Profile.Phone = updateDetails.Phone
		user.Profile.Address = updateDetails.Address
	}

	result := initializers.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&user).Error
	if result != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": result.Error()})
		return
	}

	c.IndentedJSON(201, gin.H{"data": updateDetails, "message": "User details updated successfully."})

}
