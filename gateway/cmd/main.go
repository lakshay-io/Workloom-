package main

import (
	"github.com/workloom/gateway/internal/config"
	"github.com/workloom/gateway/internal/router"
	"github.com/workloom/shared/db"
)

func main() {
	db.Init()
	config.LoadEnv()
	r := router.SetupRouter()

	port := config.GetEnv("GATEWAY_PORT", "8080")
	r.Run(":" + port)
}
