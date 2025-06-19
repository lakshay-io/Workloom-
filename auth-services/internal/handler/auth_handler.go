package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Login through google successful"})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Token is valid"})
}
