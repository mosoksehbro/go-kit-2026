package repository

import (
	"context"
	"go-kit-2026/internal/app/domain/entity"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	CreatedAt(ctx context.Context, user *entity.User) error
	UpdatedAt(ctx context.Context, user *entity.User) error
}
