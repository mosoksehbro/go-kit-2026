package service

import (
	"context"

	"go-kit-2026/internal/app/domain/entity"
)

type AuthService interface {
	Register(ctx context.Context, name, email, password string) (*entity.User, error)
	Login(ctx context.Context, email, password string) (*entity.User, string, string, error)
	Refresh(ctx context.Context, refreshToken string) (string, error)
	Logout(ctx context.Context, userID int64) error
}
