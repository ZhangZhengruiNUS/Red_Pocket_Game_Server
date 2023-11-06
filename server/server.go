package server

import (
	db "Red_Pocket_Game_Server/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Module route register
type Module interface {
	RegisterRoutes(server *Server, router *gin.Engine)
}

// Serves HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	// Initialize server
	server := &Server{
		store:  store,
		router: gin.Default(),
	}

	// Add middleware
	addMiddlewares(server.router)

	// Add routes
	var modules []Module = []Module{
		&CatalogModule{},
		&GameModule{},
		&ManageModule{},
		&UserModule{},
		&WarehouseModule{},
	}
	for _, module := range modules {
		module.RegisterRoutes(server, server.router)
	}

	// Return server
	return server
}

// Start HTTP server
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
