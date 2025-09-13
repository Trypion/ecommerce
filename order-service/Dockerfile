FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum da raiz
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY proto/ ./proto/
COPY order-service/ ./order-service/

# Build do serviço
RUN go build -o order-service-bin ./order-service/cmd/main.go

# Stage final
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/order-service-bin ./order-service

EXPOSE 50051

CMD ["./order-service"]
