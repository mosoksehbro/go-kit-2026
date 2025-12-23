package middleware

import (
	"time"

	"go-kit-2026/internal/app/utils"

	"github.com/gin-gonic/gin"
)

func RequestLogger(logger *utils.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		reqID, _ := c.Get(RequestIDKey)

		logger.Info("http_request", map[string]interface{}{
			"request_id": reqID,
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     status,
			"latency":    latency.String(),
		})
	}
}
