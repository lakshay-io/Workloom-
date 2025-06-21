package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/workloom/auth-services/internal/service"
)

func GoogleLogin(c *gin.Context) {
	service.GoogleLogin(c)
}

func GoogleAuthCallback(c *gin.Context) {
	service.GoogleCallback(c)
}

func Register(c *gin.Context) {
	service.Register(c)
}

func Login (c *gin.Context) {
	service.Login(c)
}

func LogOut(c *gin.Context) {
	service.LogOut(c)
}

func Validate(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}
