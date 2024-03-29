package main

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{},
		&models.Profile{},
		&models.Author{},
		&models.Book{},
		&models.Course{},
		&models.Student{},
	)

	if err != nil {
		panic("Failed to perform migrations: " + err.Error())
	}

	fmt.Println("Migration Success")
}
