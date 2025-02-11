// Package server contains the logic for setting up and running an HTTP server.
// Includes route handling, middleware setup, and server configuration.
package server

// Config holds configuration values for the API server, such as host and port.
type Config struct {
	Host string `envconfig:"HOST" default:"localhost"`
	Port string `envconfig:"PORT" default:"8080"`
}
