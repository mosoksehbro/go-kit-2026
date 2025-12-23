package bootstrap

import (
	"go-kit-2026/internal/app/config"
	"go-kit-2026/internal/app/middleware"
	"go-kit-2026/internal/app/utils"

	"github.com/gin-gonic/gin"
)

func NewApp(cfg *config.Config) *gin.Engine {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	app := gin.New()

	// init logger (WAJIB)
	logger := utils.NewLogger()

	// middleware order SANGAT PENTING
	app.Use(
		middleware.RequestId(),
		middleware.RequestLogger(logger),
		middleware.ErrorLogger(logger),
		middleware.Recovery(logger),
	)
	return app
}
