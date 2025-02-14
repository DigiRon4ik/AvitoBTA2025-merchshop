// Package server contains the logic for setting up and running an HTTP server.
// Includes route handling, middleware setup, and server configuration.
package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"merchshop/internal/server/handlers"
)

// Config holds configuration values for the API server, such as host and port.
type Config struct {
	Host string `envconfig:"HOST" default:"localhost"`
	Port string `envconfig:"PORT" default:"8080"`
}

type tokenManager interface {
	ParseClaims(string) (*jwt.MapClaims, error)
}

// APIServer represents the API server, including configuration, router, and services.
type APIServer struct {
	router      *gin.Engine            // HTTP router for handling requests.
	cfg         *Config                // Configuration for server settings.
	ctx         context.Context        // Application context.
	tknMng      tokenManager           // JWT Token Manager for token parsing
	usrHandlers *handlers.UserHandlers // Main handlers for user
}

// New creates a new instance of APIServer with the provided context, configuration, and services.
func New(ctx context.Context, cfg *Config, usrHandlers *handlers.UserHandlers, tknMng tokenManager) *APIServer {
	router := gin.Default()

	return &APIServer{
		router:      router,
		cfg:         cfg,
		ctx:         ctx,
		usrHandlers: usrHandlers,
		tknMng:      tknMng,
	}
}

// Start begins the HTTP server, listening on the configured host and port.
func (as *APIServer) Start() error {
	as.configureRouter() // Configure the HTTP routes
	server := &http.Server{
		Addr:         as.cfg.Host + ":" + as.cfg.Port,
		Handler:      as.router,        // Apply middleware to the router
		ReadTimeout:  time.Second * 50, // Request read timeout
		WriteTimeout: time.Second * 50, // Response Record Timeout
		IdleTimeout:  time.Second * 60, // Keep-alive connections timeout
	}
	return server.ListenAndServe() // Start the HTTP server
}
