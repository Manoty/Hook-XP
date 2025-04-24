package routes

import (
	"StreefySherehes/controllers"
	"StreefySherehes/middlewares"
	"StreefySherehes/infra"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"StreefySherehes/repositories"
	"StreefySherehes/services"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()


	//auth route setup
	userRepo := repositories.NewUserRepository()
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Replace with your Redis server address
	})

	// Initialize OTP sender
	otpSender := infra.NewOTPSender() // Replace with the actual initialization of infra.OTPSender

	authService := services.NewAuthService(userRepo, redisClient, otpSender)
	// Initialize the auth controller
	authController := controllers.NewAuthController(authService)


	// protected routess
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	{
		// event routes
	protected.POST("/events", controllers.CreateEvent)
	protected.GET("/events", controllers.GetAllEvents)
	protected.GET("/events/:id", controllers.GetEventByID)
	protected.PUT("/events/:id", controllers.UpdateEvent)
	protected.DELETE("/events/:id", controllers.DeleteEvent)




	// change password route
	protected.PUT("/changePassword", authController.ChangePassword)

	}

	


	


	r.POST("/signup", authController.SignUp)
	r.POST("/login", authController.Login)
	r.POST("/send-otp", authController.SendOTP)
	r.POST("/verify-otp", authController.VerifyOTP)


	
	

	return r

}