package routes

import (
	"github.com/gin-gonic/gin"
	"StreefySherehes/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/events", controllers.CreateEvent)
	r.GET("/events", controllers.GetAllEvents)
	

	return r

}