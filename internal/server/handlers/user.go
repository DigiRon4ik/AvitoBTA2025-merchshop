package handlers

import "github.com/gin-gonic/gin"

type UserHandlers struct {
}

func NewUserHandlers() *UserHandlers {
	return &UserHandlers{}
}

func (uh *UserHandlers) AuthHandler(c *gin.Context) {
	panic("implement me")
}

func (uh *UserHandlers) InfoHandler(c *gin.Context) {
	panic("implement me")
}

func (uh *UserHandlers) SendCoinsHandler(c *gin.Context) {
	panic("implement me")
}

func (uh *UserHandlers) BuyItemHandler(c *gin.Context) {
	panic("implement me")
}
