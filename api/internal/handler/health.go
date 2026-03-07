package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type HealthHandler struct {
	postgresPool *pgxpool.Pool
	redisClient  *redis.Client
}

func NewHealthHandler(postgresPool *pgxpool.Pool, redisClient *redis.Client) *HealthHandler {
	return &HealthHandler{
		postgresPool: postgresPool,
		redisClient:  redisClient,
	}
}

func (h *HealthHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Nexus API está online",
	})
}

func (h *HealthHandler) Health(c *gin.Context) {
	postgresStatus := h.postgresPool.Ping(c)
	redisStatus := h.redisClient.Ping(c)

	if postgresStatus != nil || redisStatus.Err() != nil {
		c.JSON(500, gin.H{
			"message":  "Nexus API está offline",
			"postgres": postgresStatus,
			"redis":    redisStatus.Err(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "Nexus API está online",
		"postgres": "OK",
		"redis":    "OK",
	})
}
