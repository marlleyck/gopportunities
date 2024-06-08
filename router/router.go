package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o router com as configurações default do gin
	router := gin.Default()

	// Define uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Run server
	router.Run() // listen and serve on 0.0.0.0:8080
}