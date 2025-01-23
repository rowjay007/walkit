package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
    "strings"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/config"
	"github.com/rowjay007/walkit/internal/middleware"
	"github.com/rowjay007/walkit/internal/routes"
	"github.com/rowjay007/walkit/pkg/logger"
	"golang.org/x/time/rate"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "github.com/rowjay007/walkit/docs" 
)

// @title           Walkit API
// @version         1.0
// @description     A fitness workout tracking API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
    logger := logger.New()

    cfg := config.LoadConfig()

    if cfg.Environment == "production" {
        gin.SetMode(gin.ReleaseMode)
    }

    router := gin.New()
    router.Use(gin.Recovery())
    router.Use(logger.GinLogger())

    // Increase rate limit for development
    limiter := middleware.NewRateLimiter(rate.Every(1*time.Second), 30) // Increased from 10 to 30
    
    // Skip rate limiting for swagger routes
    router.Use(func(c *gin.Context) {
        if strings.HasPrefix(c.Request.URL.Path, "/swagger/") {
            c.Next()
            return
        }
        middleware.RateLimiter(limiter)(c)
    })

    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = cfg.CORSAllowedOrigins
    corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
    corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
    router.Use(cors.New(corsConfig))

    // Add swagger route
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
