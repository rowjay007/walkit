package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/routes"
)

func main() {
	router := gin.Default()

	// Load routes
	routes.LoadRoutes(router)

	// Start server
	log.Println("Starting server on port :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
