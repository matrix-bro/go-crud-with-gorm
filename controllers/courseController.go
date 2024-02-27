package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"example/go-crud/serializers"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {
	var courseSerializer serializers.CourseSerializer

	if err := c.ShouldBindJSON(&courseSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Request Data"})
		return
	}

	course := models.Course{Name: courseSerializer.Name}

	result := initializers.DB.Create(&course).Error

	if result != nil {
		c.IndentedJSON(400, gin.H{"message": "Error creating new course."})
		return
	}

	c.IndentedJSON(201, gin.H{"data": courseSerializer, "message": "Created new course successfully."})
}
