package server

import (
	"github.com/gin-gonic/gin"
	db "github.com/raymond-design/transact/db/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	server.router = router
	return server
}
