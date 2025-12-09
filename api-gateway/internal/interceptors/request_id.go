package interceptors

import (
	"context"

	"github.com/Trypion/ecommerce/api-gateway/internal/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// RequestIDUnaryClientInterceptor injeta o request ID em todas as chamadas gRPC
func RequestIDUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply any,
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		// Extrai o request ID do contexto
		requestID := middleware.GetRequestID(ctx)
		if requestID != "" {
			// Injeta no metadata do gRPC
			ctx = metadata.AppendToOutgoingContext(ctx, "x-request-id", requestID)
		}

		// Continua com a chamada original
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
