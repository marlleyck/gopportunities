package router

import (
	"github.com/gin-gonic/gin"
	"github.com/marlleyck/gopportunities/handler"
)

func initializeRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings", handler.ListOpeningsHandler)

		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.PUT("/opening", handler.UpdateOpeningHandler)

		v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}