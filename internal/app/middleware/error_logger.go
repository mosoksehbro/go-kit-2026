package middleware

import (
	"go-kit-2026/internal/app/utils"

	"github.com/gin-gonic/gin"
)

func ErrorLogger(logger *utils.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				logger.Error("request_error", map[string]interface{}{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
					"error":  e.Err.Error(),
				})
			}
		}
	}
}
