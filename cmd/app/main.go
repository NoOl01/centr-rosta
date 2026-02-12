package main

import (
	"centr_rosta/internal/config"
	"centr_rosta/internal/consts"
	"centr_rosta/internal/handler"
	"centr_rosta/internal/repository"
	auth2 "centr_rosta/internal/repository/auth"
	"centr_rosta/internal/service/auth"
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
)

func main() {
	config.LoadEnv()

	logger.InitLogger()
	defer velog.Stop()

	logger.Log.Info(consts.Server, "server starting...")

	db := repository.Connect()
	repo := auth2.NewRepository(db)
	srv := auth.NewService(repo)
	h := handler.NewHandler(srv)

	if !config.Env.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE"},
		AllowHeaders:    []string{"Origin", "ContentType", "Authorization"},
	}))

	h.Router(r)

	// todo: Web build

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Env.ServerPort),
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Log.Info(consts.Server, "server started on: "+config.Env.ServerPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Log.Error(consts.Server, "server start error: "+err.Error())
		}
	}()

	<-quit
	logger.Log.Info(consts.Server, "server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Log.Error(consts.Server, "server shutdown error: "+err.Error())
	} else {
		logger.Log.Info(consts.Server, "server shutdown successfully")
	}
}
