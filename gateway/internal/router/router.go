package router

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/workloom/gateway/internal/middleware"
	"github.com/workloom/gateway/internal/proxy"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authProxy, _ := proxy.NewReverseProxy("http://localhost:8081")

	r.Any("/auth/*path", func(c *gin.Context) {
		c.Request.URL.Path = "/auth" + c.Param("path")
		authProxy.ServeHTTP(c.Writer, c.Request)
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "API Gateway is running"})
	})

	return r
}
