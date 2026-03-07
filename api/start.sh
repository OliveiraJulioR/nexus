#!/bin/sh

set -e

echo "Rodando as migrations do banco de dados..."
# Usamos as variáveis injetadas pelo docker-compose para montar a URL
migrate -path /app/migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

echo "Iniciando a API do Nexus com Air..."
exec air