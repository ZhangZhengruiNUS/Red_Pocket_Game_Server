package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Add middlewares
func addMiddlewares(router *gin.Engine) {
	router.Use(corsMiddleware())
	router.Use(loggerMiddleware())
}

// CORS middleware
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Logger middleware
func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log befre request
		path := c.Request.URL.Path
		method := c.Request.Method
		log.Printf(">> Request Received [%s]-[%s]", method, path)

		// Process request
		c.Next()

		// Log after send Response
		statusCode := c.Writer.Status()
		log.Printf("<< Response sent with status [%d]", statusCode)
	}
}
