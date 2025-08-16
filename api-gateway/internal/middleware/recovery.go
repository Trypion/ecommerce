package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			logrus.WithFields(logrus.Fields{
				"error":      err,
				"path":       c.Request.URL.Path,
				"method":     c.Request.Method,
				"request_id": c.GetString("request_id"),
			}).Error("Panic recovered")
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal_server_error",
			"message": "Something went wrong",
			"code":    http.StatusInternalServerError,
		})
	})
}
