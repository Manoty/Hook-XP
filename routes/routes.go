package routes

import (
	"StreefySherehes/controllers"
	"StreefySherehes/middlewares"

	"github.com/gin-gonic/gin"

	"StreefySherehes/repositories"
	"StreefySherehes/services"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()


	//auth route setup
	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo)
	// Initialize the auth controller
	authController := controllers.NewAuthController(authService)


	// protected routes
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


	
	

	return r

}