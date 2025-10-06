package server

import (
	"fmt"
	"gothtest/internal/config"
	"net/http"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *http.Server {
	s := &Server{
		config: cfg,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.config.Server.Port),
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
