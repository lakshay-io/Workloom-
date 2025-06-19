package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/workloom/gateway/internal/middleware"
	"github.com/workloom/gateway/internal/proxy"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())
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
