package repository

import (
	"context"
	"go-kit-2026/internal/app/domain/entity"
)

type RefreshTokenRepository interface {
	Save(ctx context.Context, refreshToken *entity.RefreshToken) error
	FindByToken(ctx context.Context, token string) (*entity.RefreshToken, error)
	DeleteByUserId(ctx context.Context, userId int64) error
}
