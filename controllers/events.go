package controllers

import (
	"StreefySherehes/databases"
	"StreefySherehes/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateEvent handles creating a new event
func CreateEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//save event to DB
	if result := databases.DB.Create(&event); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}						
	c.JSON(http.StatusOK, gin.H{"message": "Event created succesfuly", "event": event})

}