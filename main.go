package main

import (
	"example/go-crud/controllers"
	"example/go-crud/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func homePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Home Page"})
}

func main() {
	r := gin.Default()
	r.GET("/", homePage)

	// user & profile
	r.POST("/user/create", controllers.CreateUser)
	r.POST("/user/:id/profile/create", controllers.CreateProfile)
	r.GET("/user/all", controllers.AllUsers)
	r.GET("/user/profile", controllers.UserProfile)
	r.GET("/user/:id", controllers.GetUserDetails)

	// author & book
	r.POST("/author", controllers.CreateAuthor)
	r.POST("/book/:id", controllers.CreateBook)
	r.GET("/authors", controllers.GetAllAuthors)
	r.GET("/books", controllers.GetAllBooks)
	r.GET("/author/:id", controllers.GetAuthorDetails)
	r.GET("/book/:id", controllers.GetBookDetails)

	// course & student
	r.POST("/course", controllers.CreateCourse)
	r.POST("/student", controllers.CreateStudent)
	r.GET("/courses", controllers.AllCourses)
	r.GET("/students", controllers.AllStudents)
	r.GET("/course/:id", controllers.GetCourseDetails)
	r.GET("/student/:id", controllers.GetStudentDetails)
	r.POST("/enroll", controllers.EnrollStudent)

	r.Run("localhost:3000")
}
