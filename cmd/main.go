package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yitech/fugle-strategy/api/v1"
	_ "github.com/yitech/fugle-strategy/docs"
	"github.com/yitech/fugle-strategy/internal/config"
	"github.com/yitech/fugle-strategy/pkg/pubsub"
)

// Custom middleware to log requests with logrus
func LogrusLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		log.WithFields(log.Fields{
			"status":     status,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
		}).Info("incoming request")
	}
}

func main() {
	// Set logrus to output JSON format for structured logging
	log.SetFormatter(&log.JSONFormatter{})

	// Define a flag for the port number with a default value of 8080
	// Define a flag for the configuration file with a default value
	configPath := flag.String("c", "config.json5", "path to the config file")
	flag.Parse()

	// Load the configuration file
	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Resource
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	// Test connection
	pong, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Error connecting to Redis:", err)
	}
	log.Info("Connected to Redis:", pong)
	defer rdb.Close()

	// Publisher
	pub := pubsub.NewPublisher(rdb)

	// Initialize the Gin engine
	rg := gin.New()

	// Use custom Logrus logging middleware and recovery middleware
	rg.Use(LogrusLogger(), gin.Recovery())

	// Define API routes
	v1rg := rg.Group("/v1")
	api.NewSystemAPI(v1rg)
	api.NewPubAPI(v1rg, pub)

	// Swagger documentation route
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server on the specified port
	serverAddress := fmt.Sprintf(":%d", cfg.Port)
	log.Infof("Starting server on port %d", cfg.Port)
	if err := rg.Run(serverAddress); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
