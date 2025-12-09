package interceptors

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

// LoggingUnaryServerInterceptor loga todas as chamadas gRPC com request ID
func LoggingUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {
		start := time.Now()

		requestID, _ := ctx.Value("request_id").(string)

		// Chama o handler
		resp, err := handler(ctx, req)

		// Log após a execução
		logrus.WithFields(logrus.Fields{
			"request_id": requestID,
			"method":     info.FullMethod,
			"duration":   time.Since(start),
			"error":      err,
		}).Info("gRPC call")

		return resp, err
	}
}
