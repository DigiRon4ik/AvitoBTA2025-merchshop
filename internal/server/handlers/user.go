package handlers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"merchshop/internal/models"
)

var ErrInDB = errors.New("something happened to the database")

type UserHandlers struct {
	ctx       context.Context
	authSrv   authService
	tknMng    tokenManager
	usrInfSrv userInfoService
	txSrv     transactionService
	buyItmSrv buyItemService
}

func NewUserHandlers(ctx context.Context,
	authSrv authService, tknMng tokenManager, usrInfSrv userInfoService,
	txSrv transactionService, buyItmSrv buyItemService) *UserHandlers {
	return &UserHandlers{
		ctx:       ctx,
		authSrv:   authSrv,
		tknMng:    tknMng,
		usrInfSrv: usrInfSrv,
		txSrv:     txSrv,
		buyItmSrv: buyItmSrv,
	}
}

func (uh *UserHandlers) AuthHandler(c *gin.Context) {
	// switch c.GetHeader("Accept") {
	// case "application/json":
	// 	// continue
	// default:
	// 	c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": "the ‘accept’ header is not application/json"})
	// 	return
	// }

	var login models.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, ok, err := uh.authSrv.GetOrRegUser(uh.ctx, login.Username, login.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	} else if ok {
		if uh.authSrv.ComparePassword(uh.ctx, user.Password, login.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
			return
		}
	}

	tokenString, err := uh.tknMng.NewToken(strconv.Itoa(user.ID), user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "token generation failure"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (uh *UserHandlers) InfoHandler(c *gin.Context) {
	// switch c.GetHeader("Accept") {
	// case "application/json":
	// 	// continue
	// default:
	// 	c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"error": "the ‘accept’ header is not application/json"})
	// 	return
	// }
	userID := c.GetInt("user_id")

	coins, err := uh.usrInfSrv.GetCoins(uh.ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	}

	inventory, err := uh.usrInfSrv.GetInventory(uh.ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	}

	coinHistory, err := uh.usrInfSrv.GetCoinHistory(uh.ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"coins":     coins,
		"inventory": inventory,
		"coinHistory": gin.H{
			"received": coinHistory.Receiving,
			"sent":     coinHistory.Sending,
		},
	})
}

func (uh *UserHandlers) SendCoinsHandler(c *gin.Context) {
	var send models.Sending
	if err := c.ShouldBindJSON(&send); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// if send.ToUser == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "`toUser` must not be empty"})
	// 	return
	// } else if send.Amount <= 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "`amount` must be positive"})
	// 	return
	// }

	recipientID, err := uh.txSrv.GetIdRecipient(uh.ctx, send.ToUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	} else if recipientID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "`toUser` is not found"})
		return
	}

	senderID := c.GetInt("user_id")
	if senderCoins, err := uh.txSrv.GetSenderCoins(uh.ctx, senderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	} else if senderCoins < send.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you don't have enough coins"})
		return
	}

	if err := uh.txSrv.SendCoinsToUser(uh.ctx, senderID, recipientID, send.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (uh *UserHandlers) BuyItemHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	itemSlug := c.Param("item")

	item, err := uh.buyItmSrv.GetItem(uh.ctx, itemSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	} else if item == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "item not found"})
		return
	}

	buyerCoins, err := uh.buyItmSrv.GetBuyerCoins(uh.ctx, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	} else if buyerCoins < item.Price {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you don't have enough coins"})
		return
	}

	if err := uh.buyItmSrv.BuyItem(uh.ctx, userID, item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInDB.Error()})
		return
	}

	c.Status(http.StatusOK)
}
