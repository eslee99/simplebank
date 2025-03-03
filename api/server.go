package api

import (
	db "github.com/eslee99/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for our banking service.
type Server struct {
	store  *db.Store // we need access to DB
	router *gin.Engine
}

// NewServer creates a new HTTP server and setuip routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)
	// server.createAccount() will return context, context includes method to read req/ write res

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
