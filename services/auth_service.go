package services

import (

	"StreefySherehes/dto"
	"StreefySherehes/infra"
	"StreefySherehes/models"
	"StreefySherehes/repositories"
	"StreefySherehes/utils"
	"errors"
	"fmt"
	"time"
	"context"
	"math/rand"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repositories.UserRepository
	otpSender infra.OTPSender 
	redis *redis.Client
}


func NewAuthService(repo repositories.UserRepository, redisClient *redis.Client, sender infra.OTPSender) *AuthService {
	return &AuthService{
		repo: repo,
		otpSender: sender,
		redis: redisClient,
	}
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
func (as *AuthService) SendOTP(email string) error {
	otpCode := utils.GenerateOTP()
	
	key := fmt.Sprintf("otp:%s", email)
	err := as.redis.Set(context.Background(), key, otpCode, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	// Send the OTP to the user's email
	return as.otpSender.SendOTP(email, otpCode)
}

func (as *AuthService) VerifyfOTP(email, inputCode string)(bool, error) {
	key := fmt.Sprintf("otp:%s", email)
	storedCode, err := as.redis.Get(context.Background(), key).Result()
	if err != nil {
		return false, fmt.Errorf("OTP expired or not found")
	}

	if inputCode != storedCode {
		return false, fmt.Errorf("invalid OTP")

	}
	//Clear OTP from redis after succesful verification
	as.redis.Del(context.Background(), key)

	return true, nil
}
func generateOTP() string{
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
