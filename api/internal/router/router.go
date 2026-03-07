package router

import (
	"github.com/OliveiraJulioR/nexus/api/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	healthHandler := handler.NewHealthHandler()

	health := r.Group("/health")
	{
		health.GET("/ping", healthHandler.Ping)
	}

	return r
}
