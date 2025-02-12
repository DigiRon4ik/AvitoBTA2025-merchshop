package server

import "github.com/gin-gonic/gin"

func (as *APIServer) configureRouter() {
	api := as.router.Group("/api")
	{
		api.POST("/auth", func(c *gin.Context) {})
		authorized := api.Group("/", jwtMiddleware())
		{
			authorized.GET("/info", func(c *gin.Context) {})
			authorized.GET("/sendCoin", func(c *gin.Context) {})
			authorized.GET("/buy/:item", func(c *gin.Context) {})
		}
	}
}

func jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) { panic("implement me") }
}
