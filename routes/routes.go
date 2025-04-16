package routes

import (
	"github.com/gin-gonic/gin"
	"StreefySherehes/controllers"
	
	"StreefySherehes/services"
	"StreefySherehes/repositories"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// event routes
	r.POST("/events", controllers.CreateEvent)
	r.GET("/events", controllers.GetAllEvents)
	r.GET("/events/:id", controllers.GetEventByID)
	r.PUT("/events/:id", controllers.UpdateEvent)
	r.DELETE("/events/:id", controllers.DeleteEvent)


	//auth route setup
	userRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(userRepo)
	// Initialize the auth controller
	authController := controllers.NewAuthController(authService)



	r.POST("/signup", authController.SignUp)
	

	return r

}