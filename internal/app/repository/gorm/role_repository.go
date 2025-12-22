package gorm

import (
	"context"
	"go-kit-2026/internal/app/domain/entity"
	domainRepo "go-kit-2026/internal/app/domain/repository"

	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) domainRepo.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindById(ctx context.Context, id int64) (*entity.Role, error) {
	var role entity.Role

	err := r.db.WithContext(ctx).
		First(&role, id).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) FindByName(ctx context.Context, name string) (*entity.Role, error) {
	var role entity.Role
	err := r.db.WithContext(ctx).
		Where("name = ?", name).
		First(&role).Error
	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (r *roleRepository) FindByAll(ctx context.Context) ([]*entity.Role, error) {
	var role []*entity.Role

	err := r.db.WithContext(ctx).
		Find(&role)
	if err.Error != nil {
		return nil, err.Error
	}

	return role, nil
}

func (r *roleRepository) CreatedAt(ctx context.Context, role *entity.Role) error {
	return r.db.WithContext(ctx).
		Create(role).Error
}

func (r *roleRepository) UpdatedAt(ctx context.Context, role *entity.Role) error {
	return r.db.WithContext(ctx).
		Save(role).Error
}

func (r *roleRepository) GetUserRoles(ctx context.Context, userID int64) ([]*entity.Role, error) {
	var roles []*entity.Role
	err := r.db.WithContext(ctx).
		Table("roles").
		Joins("JOIN user_roles ur ON ur.role_id = roles.id").
		Where("ur.user_id = ?", userID).
		Find(&roles).Error
	return roles, err
}
