package models

import (


	"gorm.io/gorm"
	"time"
)



type User struct {
	gorm.Model
	
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	Is2FAEnabled bool 
	OTP string
	OTPExpiry time.Time
	
}

type UserInput struct {
	Username string `json:"username" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordInput struct {
	Email string `json:"email" binding:"required"`
	OldPassword string `json:"old_Password" binding:"required"`
	NewPassword string `json:"new_Password" binding:"required"`
}
type LoginResponse struct {
	User *User `json:"user"`
	Token string `json:"token"`
}