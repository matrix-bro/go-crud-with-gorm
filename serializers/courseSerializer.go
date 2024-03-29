package serializers

type CourseSerializer struct {
	Name string `json:"name" binding:"required,min=2"`
}

type StudentSerializer struct {
	CourseID  uint   `json:"course_id" binding:"required"`
	FirstName string `json:"first_name" binding:"required,min=2"`
	LastName  string `json:"last_name" binding:"required,min=2"`
}

type AllStudentsSerializer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CourseDetailsSerializer struct {
	Name     string                  `json:"name"`
	Students []AllStudentsSerializer `json:"students"`
}

type StudentDetailsSerializer struct {
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	Courses   []CourseSerializer `json:"courses"`
}

type EnrollStudentSerializer struct {
	CourseID  uint `json:"course_id" binding:"required"`
	StudentID uint `json:"student_id" binding:"required"`
}
