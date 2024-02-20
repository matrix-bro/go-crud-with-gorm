package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Profile  Profile
}

type Profile struct {
	gorm.Model
	FirstName string
	LastName  string
	UserID    uint
}
