package gorm

import (
	"context"

	"go-kit-2026/internal/app/domain/entity"
	domainRepo "go-kit-2026/internal/app/domain/repository"

	"gorm.io/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) domainRepo.PermissionRepository {
	return &permissionRepository{db: db}
}

func (r *permissionRepository) FindById(ctx context.Context, id int64) (*entity.Permission, error) {
	var permission entity.Permission

	err := r.db.WithContext(ctx).
		First(&permission, id).Error
	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *permissionRepository) FindByName(ctx context.Context, name string) (*entity.Permission, error) {
	var permission entity.Permission

	err := r.db.WithContext(ctx).
		Where("name = ?", name).
		First(&permission).Error
	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *permissionRepository) FindByAll(ctx context.Context) ([]*entity.Permission, error) {
	var permissions []*entity.Permission

	err := r.db.WithContext(ctx).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *permissionRepository) GetUserPermissions(ctx context.Context, userID int64) ([]*entity.Permission, error) {
	var perms []*entity.Permission
	err := r.db.WithContext(ctx).
		Table("permissions").
		Joins("JOIN role_permissions rp ON rp.permission_id = permissions.id").
		Joins("JOIN user_roles ur ON ur.role_id = rp.role_id").
		Where("ur.user_id = ?", userID).
		Find(&perms).Error
	return perms, err
}
