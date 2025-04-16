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
	r.GET("/events/:id", controllers.GetEventByID)
	r.PUT("/events/:id", controllers.UpdateEvent)
	r.DELETE("/events/:id", controllers.DeleteEvent)
	

	return r

}