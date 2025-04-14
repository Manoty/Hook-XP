package tests


import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"StreefySherehes/controllers"
)

func TestCreateEventHandler(t *testing.T) {
	router := gin.Default()
	router.POST("/events", controllers.CreateEvent)

	payload := `{
	"title": "Pool Party",
	"description": "A fun pool party with friends.",
	"location": "123 Pool St, Miami, FL",
	"date": "2023-10-01",
	"ticket_price": 20.00
	}`

	req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer([]byte(payload)))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	resp v = httptest.NewRecorder()
	router.ServeHTTP(resp, req)


	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "Event created successfully")


}
