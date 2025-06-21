package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/workloom/auth-services/internal/config"
	"github.com/workloom/auth-services/internal/router"
)

func main() {
	config.LoadEnv()

	port := config.GetEnv("AUTH_SERVICE_PORT","8081")

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.SetupRoutes(r)

	log.Printf("Auth service running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to run auth service: %v", err)
	}
}
