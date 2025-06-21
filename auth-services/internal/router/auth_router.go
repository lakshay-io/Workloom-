package router

import (
	"github.com/gin-gonic/gin"
	"github.com/workloom/auth-services/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.GET("google/login", handler.GoogleLogin)
		authGroup.GET("google/callback", handler.GoogleAuthCallback)
		authGroup.POST("user/register", handler.Register)
		authGroup.POST("user/login", handler.Login)
		authGroup.POST("user/logout", handler.LogOut)
		authGroup.GET("user/validate", handler.Validate)
	}
}
