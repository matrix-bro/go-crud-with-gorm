package serializers

type CourseSerializer struct {
	Name string `json:"name" binding:"required,min=2"`
}

type StudentSerializer struct {
	CourseID  uint   `json:"course_id" binding:"required"`
	FirstName string `json:"first_name" binding:"required,min=2"`
	LastName  string `json:"last_name" binding:"required,min=2"`
}