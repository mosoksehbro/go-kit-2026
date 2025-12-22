package repository

import (
	"context"
	"go-kit-2026/internal/app/domain/entity"
)

type PermissionRepository interface {
	FindById(ctx context.Context, id int64) (*entity.Permission, error)
	FindByName(ctx context.Context, name string) (*entity.Permission, error)
	FindByAll(ctx context.Context) ([]*entity.Permission, error)
	GetUserPermissions(ctx context.Context, userID int64) ([]*entity.Permission, error)
}
