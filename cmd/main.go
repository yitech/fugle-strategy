package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yitech/fugle-strategy/api/v1"
	_ "github.com/yitech/fugle-strategy/docs"
	"github.com/yitech/fugle-strategy/internal/config"
)

func main() {
	// Define a flag for the port number with a default value of 8080
	// Define a flag for the configuration file with a default value
	configPath := flag.String("c", "config.json5", "path to the config file")
	flag.Parse()
	// Load the configuration file
	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Allocate resources

	// Initialize the Gin engine
	rg := gin.Default()

	// Define a simple route
	rg.Use(gin.Logger())
	v1rg := rg.Group("/v1", gin.Recovery())
	api.NewSystemAPI(v1rg)

	// start server
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Run the server on the specified port
	serverAddress := fmt.Sprintf(":%d", cfg.Port)

	log.Printf("Starting server on port %d", cfg.Port)
	if err := rg.Run(serverAddress); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
