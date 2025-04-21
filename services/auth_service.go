package services

import (
	"StreefySherehes/models"
	"StreefySherehes/dto"
	"StreefySherehes/repositories"
	"StreefySherehes/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repositories.UserRepository
}


func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) RegisterUser(input models.UserInput) error {
	// Hash the password before saving to the database
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return err

	}
	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}
	return as.repo.CreateUser(&user)

}
func (as *AuthService) LoginUser(input models.LoginInput) (*dto.LoginResponse, error) {
	user, err := as.repo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	//compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	//generate JWT token after successful login
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dto.LoginResponse{
		Message: "Login successful",

	User: dto.UserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	},
	Token: token,
	}, nil
}

func (as *AuthService) ChangePassword(input models.ChangePasswordInput) error {
	//retrieve user by email
	user, err := as.repo.GetUserByEmail(input.Email)
	if err != nil {
		return errors.New("user not found")
	}
	//verrify old password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.OldPassword)); err != nil {
		return errors.New("incorrect old Password")
	}
	//hash the passsword
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash new password")
	}

	//update password
	user.Password = string(hashedPassword)
	return as.repo.UpdatePassword(user)

}
