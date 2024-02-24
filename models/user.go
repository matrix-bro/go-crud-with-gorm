package models

import "gorm.io/gorm"

// One to One Relationship

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Profile   *Profile
}

type Profile struct {
	gorm.Model
	UserID  uint `gorm:"uniqueIndex"`
	Phone   string
	Address string
	User    *User // Defining inverse relationship with User
}
