package handler

import (
	appErr "go-kit-2026/internal/app/error"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	e := appErr.Map(err)

	c.JSON(e.Status, gin.H{
		"success": false,
		"message": e.Message,
		"error": gin.H{
			"code": e.Code,
		},
	})
}
