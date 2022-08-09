package main

import (
	"context"
	"flag"
	"fmt"
	"go-boilerplate/app/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

type AppConfig struct {
	Env  string
	Port int
}

func main() {
	appConfig := AppConfig{
		Env: *flag.String("e", "d", "environment variable to use for running application"),
		// default or from config file
		Port: *flag.Int("p", 8080, "port to run application on"),
	}
	utils.InitializeLogger()
	utils.Logger().Info("data:", appConfig)

	router := gin.Default()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(ginzap.Ginzap(utils.FastLogger(), time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(ginzap.RecoveryWithZap(utils.FastLogger(), true))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", appConfig.Port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			utils.Logger().Fatal("unable to start server", err.Error())
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt)
	<-quit
	utils.Logger().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.Logger().Info("Server Shutdown: ", err.Error())
	}
	log.Println("Server exiting")
}
