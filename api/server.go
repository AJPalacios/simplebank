package api

import (
	db "github.com/devspace/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server struct represents the HTTP server and contains the store and router
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// Setup a new HTTP server and set up routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server
}

// Start the HTTP server on the specified address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
