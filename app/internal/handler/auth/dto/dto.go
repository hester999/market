package dto

type LogoutRequestDTO struct {
	UserID string `json:"user_id"`
}

type LoginRequestDTO struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"my_secure_password"`
}

type LoginResponseDTO struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}
