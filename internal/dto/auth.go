package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string       `json:"auth_token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	Id       int    `json:"user_id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
