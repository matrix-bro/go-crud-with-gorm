package serializers

type CourseSerializer struct {
	Name string `json:"name" binding:"required,min=2"`
}
