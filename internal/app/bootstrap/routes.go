package bootstrap

import (
	"go-kit-2026/internal/app/config"
	gormRepo "go-kit-2026/internal/app/repository/gorm"
	service "go-kit-2026/internal/app/service"

	"gorm.io/gorm"
)

type Services struct {
	Auth          service.AuthService
	Authorization service.AuthorizationService
}

func InitServices(db *gorm.DB, cfg *config.Config) *Services {
	userRepo := gormRepo.NewUserRepository(db)
	roleRepo := gormRepo.NewRoleRepository(db)
	permissionRepo := gormRepo.NewPermissionRepository(db)
	refreshRepo := gormRepo.NewRefreshTokenRepository(db)

	authService := service.NewAuthService(
		db,
		userRepo,
		refreshRepo,
		cfg.JWT.Secret,
		cfg.JWT.AccessExpire,
		cfg.JWT.RefreshExpire,
	)

	authzService := service.NewAuthorizationService(
		roleRepo,
		permissionRepo,
	)

	return &Services{
		Auth:          authService,
		Authorization: authzService,
	}
}
