package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verifica se já existe um Request ID no header
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			// Gera um novo UUID se não existir
			requestID = uuid.New().String()
		}

		// Adiciona o Request ID no contexto e response header
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)

		c.Next()
	}
}
