package main

import (
	
	"StreefySherehes/databases"
	
	"StreefySherehes/routes"
	

	"fmt"
)

func main() {
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