package models

import "gorm.io/gorm"

// Many to Many Relationship
type Course struct {
	gorm.Model
	Name     string
	Students []*Student `gorm:"many2many:course_student;"`
}

type Student struct {
	gorm.Model
	FirstName string
	LastName  string
	Courses   []*Course `gorm:"many2many:course_student;"`
}
