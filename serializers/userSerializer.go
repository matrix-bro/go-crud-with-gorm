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
