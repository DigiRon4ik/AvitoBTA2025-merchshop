// Package server contains the logic for setting up and running an HTTP server.
// Includes route handling, middleware setup, and server configuration.
package server

import (
	"github.com/gin-gonic/gin"

	"merchshop/internal/server/middlewares"
)

// configureRouter sets up the HTTP route handlers.
func (as *APIServer) configureRouter() {
	api := as.router.Group("/api")
	{
		api.POST("/auth", func(c *gin.Context) {})
		mdlwrs := middlewares.NewMiddlewares(as.tknMng)
		authorized := api.Group("/", mdlwrs.JWTMiddleware())
		{
			authorized.GET("/info", func(c *gin.Context) {})
			authorized.GET("/sendCoin", func(c *gin.Context) {})
			authorized.GET("/buy/:item", func(c *gin.Context) {})
		}
	}
}
