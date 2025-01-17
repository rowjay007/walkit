package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/config"
	"github.com/rowjay007/walkit/internal/middleware"
	"github.com/rowjay007/walkit/internal/routes"
	"github.com/rowjay007/walkit/pkg/logger"
	"golang.org/x/time/rate"
)

func main() {
    logger := logger.New()

    cfg := config.LoadConfig()

    if cfg.Environment == "production" {
        gin.SetMode(gin.ReleaseMode)
    }

    router := gin.New()

    router.Use(gin.Recovery())
    router.Use(logger.GinLogger())

    limiter := middleware.NewRateLimiter(rate.Every(1*time.Second), 10) 
    router.Use(middleware.RateLimiter(limiter))

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = cfg.CORSAllowedOrigins
    corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
    corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
    router.Use(cors.New(corsConfig))

    routes.LoadRoutes(router)

    srv := &http.Server{
        Addr:         ":" + cfg.Port,
        Handler:      router,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    go func() {
        logger.Info("Starting server on port " + cfg.Port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            logger.Fatal("Failed to start server: " + err.Error())
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    logger.Info("Shutting down server...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        logger.Fatal("Server forced to shutdown: " + err.Error())
    }

    logger.Info("Server exiting")
}
