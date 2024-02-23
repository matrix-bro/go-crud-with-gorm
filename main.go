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
	r.POST("/user/create", controllers.CreateUser)
	r.POST("/user/:id/profile/create", controllers.CreateProfile)
	r.GET("/user/all", controllers.AllUsers)
	r.GET("/user/profile", controllers.UserProfile)
	r.GET("/user/:id", controllers.GetUserById)

	r.Run("localhost:3000")
}
