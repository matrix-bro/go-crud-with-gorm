package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"example/go-crud/serializers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context) {
	var userSerializer serializers.UserSerializer

	if err := c.ShouldBindJSON(&userSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"error": "Invalid Request Data"})
		return
	}

	user := models.User{FirstName: userSerializer.FirstName, LastName: userSerializer.LastName}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"error": "Error creating new user"})
		return
	}

	c.IndentedJSON(201, gin.H{"data": userSerializer, "message": "User created successfully."})
}

func CreateProfile(c *gin.Context) {
	var user models.User
	err := CheckByID(c.Param("id"), &user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var profileSerializer serializers.ProfileSerializer
	if err := c.ShouldBindJSON(&profileSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Data"})
		return
	}

	user_profile := models.Profile{Phone: profileSerializer.Phone, Address: profileSerializer.Address, UserID: user.ID}

	profile := initializers.DB.Create(&user_profile)

	if profile.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error creating Profile"})
		return
	}

	c.IndentedJSON(201, gin.H{"data": profileSerializer, "message": "User Profile created successfully."})
}

func AllUsers(c *gin.Context) {
	var users []models.User
	var response []serializers.UserSerializer

	result := initializers.DB.Find(&users).Scan(&response)

	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error retrieving users"})
		return
	}

	c.IndentedJSON(200, gin.H{"data": response, "message": "All Users retrieved successfully."})
}

func UserProfile(c *gin.Context) {
	var userprofiles []serializers.UserProfileResponse

	result := initializers.DB.Model(&models.User{}).Select("first_name, last_name, phone, address").
		Joins("LEFT JOIN profiles on profiles.user_id=users.id").
		Scan(&userprofiles)

	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error retrieving user profiles"})
		return
	}

	c.IndentedJSON(200, gin.H{"data": userprofiles, "message": "All User profiles retrieved successfully."})
}

func GetUserDetails(c *gin.Context) {
	var user models.User
	err := CheckByID(c.Param("id"), &user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var userProfile serializers.UserProfileResponse

	result := initializers.DB.Model(&models.User{}).
		Joins("LEFT JOIN profiles on users.id = profiles.user_id").
		Select("first_name, last_name, phone, address").
		Where("users.id = ?", user.ID).Scan(&userProfile).Error

	if result != nil {
		c.IndentedJSON(400, gin.H{"message": "Error retrieving user details"})
		return
	}

	c.IndentedJSON(200, gin.H{"data": userProfile, "message": "User details retrieved successfully."})
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

	c.IndentedJSON(200, gin.H{"data": updateDetails, "message": "User details updated successfully."})

}

func DeleteUser(c *gin.Context) {
	var user models.User
	err := CheckByID(c.Param("id"), &user)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	err = initializers.DB.Select("Profile").Delete(&user).Error
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User deleted successfully."})
}
