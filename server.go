package main

import (
	"context"
	"fmt"
	"net/http"
	"odisha_gov_be/utils"
	"os"
	"os/signal"
	"syscall"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AppConfig struct {
	Env  string
	Port int
}

var (
	config *AppConfig
	server *gin.Engine
)

func InitServer(configuration *AppConfig) *gin.Engine {
	config = configuration
	utils.FastLogger().Info("initlizing server...")
	if config.Env == "p" {
		gin.SetMode(gin.ReleaseMode)
	}
	server = gin.New()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	server.Use(ginzap.Ginzap(utils.FastLogger(), time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	server.Use(ginzap.RecoveryWithZap(utils.FastLogger(), true))
	return server
}

func StartServer() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: server,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			utils.FastLogger().Fatal("unable to start server", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.FastLogger().Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		utils.FastLogger().Fatal("Server Shutdown: ", zap.Error(err))
	}
	utils.FastLogger().Info("Server Shutdown Successfull ...")
}
