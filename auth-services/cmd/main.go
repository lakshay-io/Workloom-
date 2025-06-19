package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/workloom/auth-services/internal/config"
	"github.com/workloom/auth-services/internal/router"
)

func main() {
	config.LoadEnv()

	port := config.GetEnv("AUTH_SERVICE_PORT","8081")

	r := gin.Default()

	router.SetupRoutes(r)

	log.Printf("Auth service running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to run auth service: %v", err)
	}
}
