package models

import "gorm.io/gorm"

// One to Many Relationship

type Author struct {
	gorm.Model
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Books     []*Book
}

type Book struct {
	gorm.Model
	AuthorID    uint
	Title       string `gorm:"not null"`
	Description string
	Author      *Author // Defining inverse relationship with User
}
