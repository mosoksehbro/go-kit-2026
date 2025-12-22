package routes

import (
	v1Controller "go-kit-2026/internal/app/controller/v1"
	middleware "go-kit-2026/internal/app/middleware"
	"go-kit-2026/internal/app/service"

	"github.com/gin-gonic/gin"
)

func RegisterV1(
	router *gin.Engine,
	authService service.AuthService,
	authzService service.AuthorizationService,
	jwtSecret string,
) {
	v1 := router.Group("/api/v1")

	// health
	v1.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// auth
	authController := v1Controller.NewAuthController(authService)
	v1.POST("/auth/register", authController.Register)
	v1.POST("/auth/login", authController.Login)

	// admin
	adminController := v1Controller.NewAdminController()
	admin := v1.Group("/admin")
	admin.Use(
		middleware.JWTAuth(jwtSecret),
		middleware.RequireRole(authzService, "admin"),
	)
	{
		admin.GET("/dashboard", adminController.Dashboard)
	}
}
