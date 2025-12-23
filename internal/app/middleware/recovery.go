package middleware

import (
	"net/http"
	"runtime/debug"

	"go-kit-2026/internal/app/utils"

	"github.com/gin-gonic/gin"
)

func Recovery(logger *utils.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				reqID, _ := c.Get(RequestIDKey)

				logger.Error("panic_recovered", map[string]interface{}{
					"request_id": reqID,
					"error":      err,
					"stack":      string(debug.Stack()),
				})

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"message": "internal server error",
				})
			}
		}()

		c.Next()
	}
}
