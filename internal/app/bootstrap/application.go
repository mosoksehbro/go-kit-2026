package bootstrap

import (
	"go-kit-2026/internal/app/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Application struct {
	App *gin.Engine
	DB  *gorm.DB
}

func NewApplication(cfg *config.Config) *Application {
	app := NewApp(cfg)
	db := NewDatabase(cfg.Database)

	return &Application{
		App: app,
		DB:  db,
	}
}
