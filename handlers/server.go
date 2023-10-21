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

	router.POST("/login", server.loginHandler)

	router.POST("/signup", server.signupHandler)

	router.GET("/catalog", server.catalogHandler)
	router.GET("/catalog/user", server.catalogUserHandler)
	router.POST("/catalog/buy", server.catalogBuyHandler)

	router.GET("/game/diff", server.gameDiffHandler)
	router.GET("/game/equip", server.gameEquipQueryHandler)
	router.POST("/game/equip", server.gameEquipUpdateHandler)
	router.POST("/game/end", server.gameEndHandler)

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
