# syntax=docker/dockerfile:1
FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum da raiz
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY proto/ ./proto/
COPY notification-service/ ./notification-service/

# Build do serviço
RUN go build -o notification-service-bin ./notification-service/cmd/main.go

# Stage final
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/notification-service-bin ./notification-service

EXPOSE 50053

CMD ["./notification-service"]
