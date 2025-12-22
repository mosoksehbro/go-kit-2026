package bootstrap

import (
	"go-kit-2026/internal/app/config"
	"go-kit-2026/internal/app/middleware"

	"github.com/gin-gonic/gin"
)

func NewApp(cfg *config.Config) *gin.Engine {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	app := gin.New()

	app.Use(gin.Logger())
	app.Use(middleware.RequestLogger())
	app.Use(gin.Recovery())

	return app
}
