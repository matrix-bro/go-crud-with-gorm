package serializers

type UserResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserProfileResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type UpdateUserDetailsSerializer struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Address   string `json:"address" binding:"required"`
}
