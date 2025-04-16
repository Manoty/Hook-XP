package services


import (
	"StreefySherehes/models"
	"StreefySherehes/repositories"
	"StreefySherehes/utils"
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
		Username : input.Username,
		Email : input.Email,
		Password : hashedPassword,
	}
	return as.repo.CreateUser(&user)

}