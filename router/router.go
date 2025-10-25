package router

import "github.com/gin-gonic/gin"

func Initializer() {
	// Initialize Router, with default Gin settings
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	// Run the server
	router.Run(":8080")
}
