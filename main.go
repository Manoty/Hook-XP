package main

import (
	"StreefySherehes/config"
	"StreefySherehes/databases"
	"StreefySherehes/infra"
	"StreefySherehes/repositories"
	"StreefySherehes/services"

	"StreefySherehes/routes"
	

	"fmt"
)

func main() {

	//connect to redis
	redisClient := config.InitRedis()

	smtpSender := &infra.SmtpSender{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "youremail@gmail.com",
		Password: "app_password",
		From:     "youremail@gmail.com",
	}
	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo, redisClient, smtpSender)

	infraSender := infra.NewOTPSender()
	authService = services.NewAuthService(userRepo, redisClient, infraSender)


	email := "user@example.com"
	if err := authService.SendOTP(email); err != nil {
		fmt.Println("failed to send otp:", err)
	}
	// Connect to the database
	databases.ConnectDatabase()

	//
	
	// Set up the router
	r := routes.SetupRoutes()

	// Start the server
	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
	
	}
}