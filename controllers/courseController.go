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

func CreateStudent(c *gin.Context) {
	var studentSerializer serializers.StudentSerializer

	if err := c.ShouldBindJSON(&studentSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Request Data"})
		return
	}

	courseId := studentSerializer.CourseID
	var course models.Course
	if err := initializers.DB.First(&course, courseId).Error; err != nil {
		c.IndentedJSON(400, gin.H{"message": "Course does not exists."})
		return
	}

	student := models.Student{
		FirstName: studentSerializer.FirstName,
		LastName:  studentSerializer.LastName,
		Courses:   []*models.Course{&course},
	}

	if err := initializers.DB.Create(&student).Error; err != nil {
		c.IndentedJSON(400, gin.H{"message": "Error creating student"})
		return
	}

	c.IndentedJSON(201, gin.H{"data": studentSerializer, "message": "Created new student successfully."})
}
