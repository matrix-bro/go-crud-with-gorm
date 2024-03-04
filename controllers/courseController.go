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

func GetStudentDetails(c *gin.Context) {
	var student models.Student
	err := CheckByID(c.Param("id"), &student)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var studentDetails serializers.StudentDetailsSerializer
	result := initializers.DB.Preload("Courses").Find(&student).Error

	if result != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": result.Error()})
		return
	}

	studentDetails.FirstName = student.FirstName
	studentDetails.LastName = student.LastName

	for _, course := range student.Courses {
		course_serializer := serializers.CourseSerializer{
			Name: course.Name,
		}

		studentDetails.Courses = append(studentDetails.Courses, course_serializer)
	}

	c.IndentedJSON(201, gin.H{"data": studentDetails, "message": "Student details retrieved successfully."})
}

func EnrollStudent(c *gin.Context) {
	// check if the student is already enrolled in the same course
	// if not then enroll
	var enrollStudent serializers.EnrollStudentSerializer

	if err := c.ShouldBindJSON(&enrollStudent); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Request Data"})
		return
	}

	var course models.Course
	err := initializers.DB.First(&course, enrollStudent.CourseID).Error
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var student models.Student
	err = initializers.DB.First(&student, enrollStudent.StudentID).Error
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	check_if_already_enrolled := initializers.DB.Model(&student).Where("course_id = ?", enrollStudent.CourseID).Association("Courses").Count()
	if check_if_already_enrolled > 0 {
		c.AbortWithStatusJSON(400, gin.H{"error": "Student is already enrolled in this course."})
		return
	}

	result := initializers.DB.Model(&student).Association("Courses").Append(&course)

	if result != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": result.Error()})
		return
	}

	c.IndentedJSON(201, gin.H{"data": enrollStudent, "message": "Student enrolled successfully."})

}

func UpdateCourseDetails(c *gin.Context) {
	var course models.Course
	err := CheckByID(c.Param("id"), &course)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var courseSerializer serializers.CourseSerializer
	if err := c.ShouldBindJSON(&courseSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Data"})
		return
	}

	course.Name = courseSerializer.Name

	err = initializers.DB.Save(&course).Error
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"data": courseSerializer, "message": "Course details updated successfully."})
}

func DeleteCourse(c *gin.Context) {
	var course models.Course
	err := CheckByID(c.Param("id"), &course)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	err = initializers.DB.Select("Students").Delete(&course).Error
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Course deleted successfully."})
}
