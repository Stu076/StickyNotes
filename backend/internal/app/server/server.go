package server

import (
	"janstupica/StickyNotes/logger"

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

func (s *Server) Run(port int) {

}
