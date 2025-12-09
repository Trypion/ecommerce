package interceptors

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RequestIDUnaryServerInterceptor extrai o request ID do metadata
func RequestIDUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		// Extrai o request ID do metadata
		var requestID string
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			if values := md.Get("x-request-id"); len(values) > 0 {
				requestID = values[0]
			}
		}

		// Gera um novo se não existir (fallback)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Adiciona ao contexto para uso nos handlers
		ctx = context.WithValue(ctx, "request_id", requestID)

		return handler(ctx, req)
	}
}
