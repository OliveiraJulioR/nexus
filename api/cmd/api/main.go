package main

import (
	"context"
	"log"

	"github.com/OliveiraJulioR/nexus/api/infra/config"
	"github.com/OliveiraJulioR/nexus/api/infra/database"
	"github.com/OliveiraJulioR/nexus/api/internal/router"
)

func main() {
	// Carrega as variáveis de ambiente
	config.LoadEnv()
	// Inicia as conexões com o banco de dados
	postgresPool, err := database.NewPostgresConnection(context.Background(), config.GetConnectionString())
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
	defer postgresPool.Close()

	redisClient, err := database.NewRedisConnection(context.Background(), config.GetRedisConnectionString())
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
	defer redisClient.Close()

	r := router.SetupRouter(postgresPool, redisClient)
	r.Run(":8080")
}
