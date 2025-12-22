package v1

import (
	"github.com/gin-gonic/gin"
)

type AdminController struct {
}

func NewAdminController() *AdminController {
	return &AdminController{}
}

func (h *AdminController) Dashboard(c *gin.Context) {
	c.JSON(200, gin.H{"message": "admin dashboard"})
}
