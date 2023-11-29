package server

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
	db  *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{gin: gin.Default(), db: db}
}

func (s *Server) Run() error {
	s.GameRoutes()
	s.gin.Run(":3000")
	return nil
}
