package gorm

import (
	"context"
	"time"

	"go-kit-2026/internal/app/domain/entity"
	domainRepo "go-kit-2026/internal/app/domain/repository"

	"gorm.io/gorm"
)

type refreshTokenRepository struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) domainRepo.RefreshTokenRepository {
	return &refreshTokenRepository{db: db}
}

func (r *refreshTokenRepository) Save(ctx context.Context, token *entity.RefreshToken) error {
	return r.db.WithContext(ctx).
		Create(token).Error
}

func (r *refreshTokenRepository) FindByToken(ctx context.Context, token string) (*entity.RefreshToken, error) {
	var refresh entity.RefreshToken

	err := r.db.WithContext(ctx).
		Where("token = ? AND expires_at > ?", token, time.Now()).
		First(&refresh).Error
	if err != nil {
		return nil, err
	}

	return &refresh, nil
}

func (r *refreshTokenRepository) DeleteByUserId(ctx context.Context, userID int64) error {
	return r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Delete(&entity.RefreshToken{}).Error
}
