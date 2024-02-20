package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Home Page"})
}

func main() {
	r := gin.Default()
	r.GET("/", homePage)

	r.Run("localhost:3000")
}
