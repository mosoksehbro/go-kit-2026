package service

import (
	"context"
	domainRepo "go-kit-2026/internal/app/domain/repository"
)

type authorizationService struct {
	roleRepo       domainRepo.RoleRepository
	permissionRepo domainRepo.PermissionRepository
}

func NewAuthorizationService(roleRepo domainRepo.RoleRepository, permissionRepo domainRepo.PermissionRepository) AuthorizationService {
	return &authorizationService{
		roleRepo:       roleRepo,
		permissionRepo: permissionRepo,
	}
}

func (s *authorizationService) HasRole(ctx context.Context, userID int64, role string) (bool, error) {
	roles, err := s.roleRepo.GetUserRoles(ctx, userID)
	if err != nil {
		return false, err
	}
	for _, r := range roles {
		if r.Name == role {
			return true, nil
		}
	}
	return false, nil
}

func (s *authorizationService) HasPermission(ctx context.Context, userID int64, perm string) (bool, error) {
	perms, err := s.permissionRepo.GetUserPermissions(ctx, userID)
	if err != nil {
		return false, err
	}
	for _, p := range perms {
		if p.Name == perm {
			return true, nil
		}
	}
	return false, nil
}
