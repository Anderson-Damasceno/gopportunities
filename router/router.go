package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Inicializa o Router usando as configurações do gin.
	r := gin.Default()

	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080")
}