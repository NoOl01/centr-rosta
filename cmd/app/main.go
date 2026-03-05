package main

import (
	"centr_rosta/internal/bootstrap"
	"centr_rosta/internal/config"
	"centr_rosta/internal/consts/log_names"
	"centr_rosta/pkg/logger"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nool01/velog/pkg/velog"
	"github.com/redis/go-redis/v9"
)

func main() {
	config.LoadEnv()

	logger.InitLogger()
	defer velog.Stop()

	logger.Log.Info(log_names.Server, "server starting...")

	rdb, h := bootstrap.Bootstrap()
	defer func(rdb *redis.Client) {
		err := rdb.Close()
		if err != nil {
			logger.Log.Error(log_names.Server, "Failed to close Redis connection")
			return
		}
	}(rdb)

	if !config.Env.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Origin", "ContentType", "Authorization", "X-Session-ID"},
	}))

	h.Router(r)

	// todo: Web build

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", config.Env.ServerPort),
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Log.Info(log_names.Server, "server started on: "+config.Env.ServerPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Error(log_names.Server, "server start error: "+err.Error())
		}
	}()

	<-quit
	logger.Log.Info(log_names.Server, "server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Error(log_names.Server, "server shutdown error: "+err.Error())
	} else {
		logger.Log.Info(log_names.Server, "server shutdown successfully")
	}
}
