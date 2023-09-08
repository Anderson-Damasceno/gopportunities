package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	_ = router.Run(":8080")
}
