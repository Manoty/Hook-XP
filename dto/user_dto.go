package dto

type UserDTO struct {
	ID uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
}
type LoginResponse struct {
	Message string `json:"message"`
	User UserDTO `json:"user"`
	Token string `json:"token"`
}
type ChangePasswordResponse struct {
	Message string `json:"message"`
	User UserDTO `json:"user"`
}
