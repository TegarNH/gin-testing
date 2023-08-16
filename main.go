package main

import (
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	// router.Use(cors.Default())

	// config := cors.DefaultConfig()
	// config.AllowWildcard = true
	// config.AllowAllOrigins = true
	// config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	// config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Token"}

	// router.Use(cors.New(config))
	// router.Use(cors.New(config))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
