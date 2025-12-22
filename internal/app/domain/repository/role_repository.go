package repository

import (
	"context"
	"go-kit-2026/internal/app/domain/entity"
)

type RoleRepository interface {
	FindById(ctx context.Context, id int64) (*entity.Role, error)
	FindByName(ctx context.Context, name string) (*entity.Role, error)
	FindByAll(ctx context.Context) ([]*entity.Role, error)
	GetUserRoles(ctx context.Context, userID int64) ([]*entity.Role, error)
	CreatedAt(ctx context.Context, role *entity.Role) error
	UpdatedAt(ctx context.Context, role *entity.Role) error
}
