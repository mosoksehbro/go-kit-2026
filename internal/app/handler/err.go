package handler

import (
	"net/http"

	"go-kit-2026/internal/app/service"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*service.AppError); ok {
		c.JSON(appErr.Status, gin.H{
			"success": false,
			"message": appErr.Message,
			"error": gin.H{
				"code": appErr.Code,
			},
		})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"success": false,
		"message": "internal server error",
	})
}
