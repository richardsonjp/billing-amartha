package routes

import (
	"github.com/gin-gonic/gin"
	"go-skeleton/internal/middlewares"
)

// Init initialize application routes
// DO NOT FORGET TO READ README#routes FIRST!
func Init(mode string) *gin.Engine {
	gin.SetMode(mode)
	r := gin.Default()
	r.HandleMethodNotAllowed = true
	r.Use(middlewares.LanguageAccept())

	// Setup pingpong
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1Route := r.Group("/api/v1")
	V1Route(v1Route)

	return r
}
