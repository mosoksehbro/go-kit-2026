package service

import (
	"context"
)

type AuthorizationService interface {
	HasRole(ctx context.Context, userID int64, role string) (bool, error)
	HasPermission(ctx context.Context, userID int64, perm string) (bool, error)
}
