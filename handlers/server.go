package handler

import (
	db "Red_Pocket_Game_Server/db/sqlc"

	"github.com/gin-gonic/gin"
)

// server serves HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.Use(corsMiddleware())

	router.POST("/login", server.loginHandler)

	router.POST("/signup", server.signupHandler)

	router.GET("/catalog", server.catalogHandler)
	router.GET("/catalog/user", server.catalogUserHandler)
	router.POST("/catalog/buy", server.catalogBuyHandler)

	router.GET("/game/diff", server.gameDiffHandler)
	router.GET("/game/equip", server.gameEquipQueryHandler)
	router.POST("/game/equip", server.gameEquipUpdateHandler)
	router.POST("/game/end", server.gameEndHandler)

	router.GET("/warehouse", server.warehouseInfoQueryHandler)
	router.GET("/warehouse/rolltable", server.warehouseRolltableQueryHandler)
	router.POST("/warehouse/rolltable", server.warehouseRolltableUpdateHandler)

	server.router = router
	return server
}

// start runs the HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// handle error response
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// handle common response
func commonResponse(msg string) gin.H {
	return gin.H{"error": msg}
}

// cors middleware
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
