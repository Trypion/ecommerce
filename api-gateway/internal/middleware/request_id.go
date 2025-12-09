package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type contextKey string

const requestIDKey contextKey = "request_id"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")

		if requestID != "" && !isValidUUID(requestID) {
			requestID = ""
		}

		// Gera um novo Request ID se não existir ou for inválido
		if requestID == "" {
			requestID = uuid.New().String()
		}

		c.Set("request_id", requestID)

		ctx := context.WithValue(c.Request.Context(), requestIDKey, requestID)
		c.Request = c.Request.WithContext(ctx)

		c.Header("X-Request-ID", requestID)

		c.Next()
	}
}

func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(requestIDKey).(string); ok {
		return requestID
	}
	return ""
}

func isValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
