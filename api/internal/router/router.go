package router

import (
	"github.com/OliveiraJulioR/nexus/api/internal/handler"
	"github.com/OliveiraJulioR/nexus/api/internal/repository/implementations"
	usecase "github.com/OliveiraJulioR/nexus/api/internal/usecase/task"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func SetupRouter(postgresPool *pgxpool.Pool, redisClient *redis.Client) *gin.Engine {
	r := gin.Default()

	taskRepository := implementations.NewTaskPostgresRepository(postgresPool)

	createTaskUseCase := usecase.NewCreateTaskUseCase(taskRepository)
	updateStatusUseCase := usecase.NewUpdateStatusUseCase(taskRepository)
	findAllUseCase := usecase.NewFindAllUseCase(taskRepository)
	deleteTaskUseCase := usecase.NewDeleteTaskUseCase(taskRepository)
	findByIdUseCase := usecase.NewFindByIDUseCase(taskRepository)
	updateTaskUseCase := usecase.NewUpdateTask(taskRepository)

	healthHandler := handler.NewHealthHandler(postgresPool, redisClient)
	taskHandler := handler.NewTaskHandler(createTaskUseCase, updateStatusUseCase, findAllUseCase, findByIdUseCase, deleteTaskUseCase, updateTaskUseCase)

	health := r.Group("/health")
	{
		health.GET("/ping", healthHandler.Ping)
		health.GET("/", healthHandler.Health)
	}

	task := r.Group("/task")
	{
		task.POST("/", taskHandler.CreateTask)
		task.PATCH("/:id/status", taskHandler.UpdateStatus)
		task.GET("/", taskHandler.FindAll)
	}

	return r
}
