FROM golang:1.25-alpine

# Necessário para pacotes Go e utilitários
RUN apk add --no-cache git make

WORKDIR /app

# Instala o Air para live reload
RUN go install github.com/air-verse/air@latest

# Instala o golang-migrate com suporte APENAS para Postgres (deixa a imagem mais leve)
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Dá permissão de execução para o script
RUN chmod +x /app/start.sh

EXPOSE 8080

# Troca o Air direto pelo nosso script
CMD ["/app/start.sh"]