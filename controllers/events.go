package controllers

import (
	"StreefySherehes/databases"
	"StreefySherehes/models"
	"net/http"
	"strconv"

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

// GetAllEvents handles retrieving all events
func GetAllEvents(c *gin.Context){
	var events [] models.Event
	if result := databases.DB.Find(&events); result.Error != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch events"})
		return

	}
	c.JSON(http.StatusOK, events)
	
}

// GetEventByID handles retrieving a single event by ID
func GetEventByID(c *gin.Context) {
	idParam := c.Param ("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		return
	}

	// Find the event by ID
	var event models.Event
	if result := databases.DB.First(&event, uint(id)); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}
	c.JSON(http.StatusOK, event)
}

// UpdateEvent handles updating an existing event
func UpdateEvent(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi((idParam))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event ID"})
		return
	}

	// Find the event by ID
	var event models.Event
	if result := databases.DB.First(&event, uint(id)); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		return
	}

	// Bind the request body to the event struct
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	// Update the event in the database
	if result := databases.DB.Save(&event); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}	
	c.JSON(http.StatusOK, gin.H{"message": "Efent updated successfuly", "event": event})
}