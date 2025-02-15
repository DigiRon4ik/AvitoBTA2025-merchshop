// Package main = entry point.
package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"merchshop/internal/config"
	"merchshop/internal/db"
	"merchshop/internal/hasher"
	"merchshop/internal/logger"
	"merchshop/internal/modules/authentication"
	"merchshop/internal/modules/buy_item"
	"merchshop/internal/modules/jwt_token_manager"
	"merchshop/internal/modules/transaction"
	"merchshop/internal/modules/user_info"
	"merchshop/internal/server"
	"merchshop/internal/server/handlers"
)

var version = "v1.1.0"

func main() {
	cfg := config.MustLoad()     // config loading
	logg := logger.Init(cfg.Log) // logger initialization
	logg.Info("Application loading...")

	ctx, ctxCancel := context.WithCancel(context.Background()) // general context

	storage, err := db.NewPostgresPool(ctx, cfg.DB) // storage creation
	if err != nil {
		logg.Error("db.NewPostgresPool", "err", err.Error())
		os.Exit(1)
	}

	passwdHasher := hasher.New()                  // hash creation
	tknMng, err := jwt_token_manager.New(cfg.JWT) // token manager creation
	if err != nil {
		logg.Error("jwt_token_manager.New", "err", err.Error())
		os.Exit(1)
	}
	authSrv := authentication.New(storage, passwdHasher) // authentication module creation
	usrInfSrv := user_info.New(storage)                  // creating a user information module
	txSrv := transaction.New(storage)                    // transaction module creation
	buyItmSrv := buy_item.New(storage)                   // creating an item purchase module

	// creating the main request handler
	usrHandlers := handlers.NewUserHandlers(ctx, authSrv, tknMng, usrInfSrv, txSrv, buyItmSrv)
	// server creation
	serv := server.New(ctx, cfg.APIServer, usrHandlers, tknMng)

	// server startup
	go func() {
		logg.Info("Application Started! " + version)
		if err = serv.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logg.Error("serv.Start", "err", err)
			os.Exit(1)
		}
	}()

	// graceful shutdown
	sigCtx, sigStop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	<-sigCtx.Done()
	sigStop()
	logg.Info("Shutting down gracefully...")

	ctxTimeOut, ctxTimeOutCancel := context.WithTimeout(ctx, 5*time.Second)
	defer ctxTimeOutCancel()

	if err = serv.Shutdown(ctxTimeOut); err != nil {
		logg.Error("serv.Shutdown", "err", err.Error())
	}

	if storage != nil {
		storage.Close()
	}

	ctxCancel()
	logg.Info("Application Stopped!")
}
