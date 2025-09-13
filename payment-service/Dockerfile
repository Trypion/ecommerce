FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum da raiz
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY proto/ ./proto/
COPY payment-service/ ./payment-service/

# Build do serviço
RUN go build -o payment-service-bin ./payment-service/cmd/main.go

# Stage final
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/payment-service-bin ./payment-service

EXPOSE 50052

CMD ["./payment-service"]
