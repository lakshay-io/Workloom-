package router

import (
	"github.com/gin-gonic/gin"
	"github.com/workloom/auth-services/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("google/login", handler.Login)
		authGroup.POST("/register", handler.Register)
		authGroup.GET("/validate", handler.Validate)
	}
}
