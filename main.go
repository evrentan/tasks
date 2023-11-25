package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/evrentan/tasks/internal/config"
)

func main() {
	cfg, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatalf("Error reading config file. Error: %v", err)
	}

	logger := config.NewLogger(cfg)
	config.GetDbConnection(cfg, logger)
	logger.Infof("Starting server on port %d", cfg.Port)

	r := gin.Default()
	r.GET("/hello", hello)
	r.GET("/ping", ping)
	log.Fatal(r.Run(fmt.Sprintf(":%d", cfg.Port)))
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
