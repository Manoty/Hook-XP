package controllers

import (
	"StreefySherehes/models"
	"net/http"
	"StreefySherehes/services"
	"StreefySherehes/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(service *services.AuthService) *AuthController {
	return &AuthController{service: service}
}

// AuthController handles authentication-related requests
func (ac *AuthController) SignUp(c *gin.Context) {
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ac.service.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Created Successfuly"})
}
func(ac *AuthController) Login(c *gin.Context) {
	var input models.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := ac.service.VerifyUserCredentials(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	err = ac.service.SendOTP(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}
	

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful. OTP sent to email, please verify to complete your login",
	})
}
	

func (ac AuthController) SendOTP(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
		UserID uint 
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err := ac.service.SendOTP(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return

	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP Sent successfully"})
}

func (ac *AuthController) VerifyOTP(c *gin.Context) {
	var request struct {
		Email string `json:"email" `
		OTP   string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	valid, err := ac.service.VerifyfOTP(request.Email, request.OTP)
	if !valid || err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return
    }
	user, err := ac.service.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}


	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "login successfully", "token": token})
}


func (ac *AuthController) ChangePassword(c *gin.Context) {
	var input models.ChangePasswordInput

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ac.service.ChangePassword(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfuly"})
}

	
