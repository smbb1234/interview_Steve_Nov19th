package main

import (
	"steveInterviewMod/config"
	"steveInterviewMod/routes"
	"steveInterviewMod/services"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize the Gin engine
	r := gin.Default()

	// Setting up routing
	routes.SetupRoutes(r)

	config.Setup()

	// Start the challenge ID updater in a separate goroutine
	go services.UpdateChallengeIDPeriodically()

	// Start the server
	r.Run(":8080")
}
