# Carrega as variáveis do .env (opcional, mas muito útil)
include .env

# Monta a URL de conexão usando as variáveis ou valores fixos para facilitar o uso local
DB_URL=postgres://nexus_user:nexus_password@localhost:5432/nexus_db?sslmode=disable

# Cria uma nova migration em branco (Ex: make migrate-create name=add_users_table)
migrate-create:
	migrate create -ext sql -dir migrations/up -seq $(name)

# Roda todas as migrations pendentes (Sobe a estrutura do banco)
migrate-up:
	migrate -path migrations/up -database "$(DB_URL)" -verbose up

# Desfaz a última migration (Útil se você errar algo no SQL)
migrate-down:
	migrate -path migrations/down -database "$(DB_URL)" -verbose down 1

# Derruba tudo (CUIDADO: apaga todas as tabelas)
migrate-force-drop:
	migrate -path migrations/down -database "$(DB_URL)" -verbose drop