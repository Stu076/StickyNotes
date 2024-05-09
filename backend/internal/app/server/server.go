package server

import (
	"janstupica/StickyNotes/logger"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	Log zerolog.Logger
}

func New() *Server {
	log := logger.Init("debug")

	return &Server{
		Log: log,
	}
}

func (s *Server) Run() error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api/v1")

	_ = api
	return nil
}
