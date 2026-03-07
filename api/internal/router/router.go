package router

import (
	"github.com/OliveiraJulioR/nexus/api/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func SetupRouter(postgresPool *pgxpool.Pool, redisClient *redis.Client) *gin.Engine {
	r := gin.Default()

	healthHandler := handler.NewHealthHandler(postgresPool, redisClient)

	health := r.Group("/health")
	{
		health.GET("/ping", healthHandler.Ping)
		health.GET("/", healthHandler.Health)
	}

	return r
}
