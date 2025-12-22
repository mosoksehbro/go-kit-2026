package middleware

import (
	service "go-kit-2026/internal/app/service"

	"github.com/gin-gonic/gin"
)

func RequirePermission(
	authz service.AuthorizationService,
	perm string,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("user_id")

		ok, err := authz.HasPermission(c.Request.Context(), userID, perm)
		if err != nil || !ok {
			c.AbortWithStatusJSON(403, gin.H{"message": "forbidden"})
			return
		}
		c.Next()
	}
}
