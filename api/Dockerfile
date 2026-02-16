FROM golang:1.25-alpine

# Necessário para pacotes Go
RUN apk add --no-cache git

WORKDIR /app

# Esta versão já está no namespace air-verse e aceita Go 1.23
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]