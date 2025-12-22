package gorm

import (
	"context"

	"go-kit-2026/internal/app/domain/entity"
	domainRepo "go-kit-2026/internal/app/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domainRepo.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).
		First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) CreatedAt(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).
		Create(user).Error
}

func (r *userRepository) UpdatedAt(ctx context.Context, user *entity.User) error {
	return r.db.WithContext(ctx).
		Save(user).Error
}
