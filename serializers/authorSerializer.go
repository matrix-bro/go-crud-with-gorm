package serializers

type AuthorSerializer struct {
	FirstName string `json:"first_name" binding:"required,min=2"`
	LastName  string `json:"last_name" binding:"required,min=2"`
}

type BookSerializer struct {
	Title       string `json:"title" binding:"required,min=2"`
	Description string `json:"description" binding:"required,min=2"`
}

type AuthorDetailsSerializer struct {
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Books     []BookSerializer `json:"books"`
}
