FROM golang:1.25.1-alpine AS builder

WORKDIR /app

# Copiar go.mod e go.sum da raiz
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar código fonte
COPY proto/ ./proto/
COPY user-service/ ./user-service/
# Build do serviço
RUN go build -o user-service-bin ./user-service/cmd/main.go

# Stage final
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/user-service-bin ./user-service

EXPOSE 50054

CMD ["./user-service"]