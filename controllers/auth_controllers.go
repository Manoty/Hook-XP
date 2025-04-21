package controllers

import (
	"StreefySherehes/models"
	"net/http"
	"StreefySherehes/services"

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
	token, err := ac.service.LoginUser(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token": token,
	})
	
}

func (ac *AuthController) ChangePassword(c *gin.Context) {
	var input models.ChangePasswordInput

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ac.service.ChangePassword(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfuly"})
}
	
