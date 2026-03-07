package handler

import "github.com/gin-gonic/gin"

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Nexus API está online",
	})
}
