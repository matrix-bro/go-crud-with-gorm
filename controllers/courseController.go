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

func AllCourses(c *gin.Context) {
	var courseSerializer []serializers.CourseSerializer
	result := initializers.DB.Model(&models.Course{}).Find(&courseSerializer).Error

	if result != nil {
		c.IndentedJSON(400, gin.H{"message": "Error retrieving all courses"})
		return
	}

	c.IndentedJSON(201, gin.H{"data": courseSerializer, "message": "All Courses retrieved successfully."})
}

func AllStudents(c *gin.Context) {
	var studentSerializer []serializers.AllStudentsSerializer
	result := initializers.DB.Model(&models.Student{}).Find(&studentSerializer).Error

	if result != nil {
		c.IndentedJSON(400, gin.H{"message": "Error retrieving all students"})
		return
	}

	c.IndentedJSON(201, gin.H{"data": studentSerializer, "message": "All Students retrieved successfully."})
}

func GetCourseDetails(c *gin.Context) {
	var course models.Course
	err := CheckByID(c.Param("id"), &course)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var courseDetails serializers.CourseDetailsSerializer
	result := initializers.DB.Preload("Students").Find(&course).Error

	if result != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": result.Error()})
		return
	}

	courseDetails.Name = course.Name

	for _, student := range course.Students {
		student_serializer := serializers.AllStudentsSerializer{
			FirstName: student.FirstName,
			LastName:  student.LastName,
		}

		courseDetails.Students = append(courseDetails.Students, student_serializer)
	}

	c.IndentedJSON(201, gin.H{"data": courseDetails, "message": "Course details retrieved successfully."})
}
