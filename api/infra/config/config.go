package config

import (
	"fmt"
	"os"
)

// Postgress Host
var DB_HOST = "localhost"
var DB_PORT = "5432"
var DB_USER = "postgres"
var DB_PASSWORD = "[PASSWORD]"
var DB_NAME = "nexus"

// Redis Host
var REDIS_HOST = "localhost"
var REDIS_PORT = "6379"
var REDIS_PASSWORD = ""

func LoadEnv() {
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")

	REDIS_HOST = os.Getenv("REDIS_HOST")
	REDIS_PORT = os.Getenv("REDIS_PORT")
	REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
}

func GetConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
}

func GetRedisConnectionString() string {
	return fmt.Sprintf("%s:%s", REDIS_HOST, REDIS_PORT)
}
