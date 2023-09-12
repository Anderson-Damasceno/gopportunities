package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	initializeRoutes(router)

	//Run the server
	_ = router.Run(":8080")
}
