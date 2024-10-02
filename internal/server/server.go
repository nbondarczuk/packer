package server

import (
	"log/slog"

	"github.com/gin-gonic/gin"

	"packer/internal/config"
	"packer/internal/handlers"
	"packer/internal/logging"
	"packer/internal/middleware"
)

// Server links handlers to paths via routes.
type Server struct {
	router *gin.Engine
}

// New creates server with gin framework.
func New() (*Server, error) {
	gin.SetMode(gin.ReleaseMode)

	s := &Server{
		router: gin.New(),
	}

	s.router.Use(middleware.ResponseLogger())
	s.router.Use(middleware.RequestLogger())

	s.RegisterHandlers()

	return s, nil
}

// Run the gin server on routes.
func (s *Server) Run() error {
	port := config.ServerHTTPPort()
	logging.Logger.Info("Starting HTTP server", slog.String("port", port))
	return s.router.Run(":" + port)
}

// RegisterHandlers links handlers to API points.
func (s *Server) RegisterHandlers() {
	s.router.POST("/pack", handlers.PackHandler)
}
