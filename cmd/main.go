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
)

func main() {
	// Define a flag for the port number with a default value of 8080
	port := flag.String("port", "8080", "port to run the server on")

	// Parse command-line flags
	flag.Parse()

	// Initialize the Gin engine
	rg := gin.Default()

	// Define a simple route
	rg.Use(gin.Logger())
	v1rg := rg.Group("/v1", gin.Recovery())
	api.NewSystemAPI(v1rg)

	// start server
	rg.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Run the server on the specified port
	serverAddress := fmt.Sprintf(":%s", *port)

	log.Printf("Starting server on port %s", *port)
	if err := rg.Run(serverAddress); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
