package service

import (
	"context"
	"time"

	"go-kit-2026/internal/app/domain/entity"
	domainRepo "go-kit-2026/internal/app/domain/repository"
	"go-kit-2026/internal/app/utils"

	"gorm.io/gorm"
)

type authService struct {
	db               *gorm.DB
	userRepo         domainRepo.UserRepository
	refreshTokenRepo domainRepo.RefreshTokenRepository
	jwtSecret        string
	accessExpire     int
	refreshExpire    int
}

func NewAuthService(db *gorm.DB, userRepo domainRepo.UserRepository, refreshTokenRepo domainRepo.RefreshTokenRepository, jwtSecret string, accessExpire int, refreshExpire int) AuthService {
	return &authService{
		db:               db,
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		jwtSecret:        jwtSecret,
		accessExpire:     accessExpire,
		refreshExpire:    refreshExpire,
	}
}

func (s *authService) Login(ctx context.Context, email, password string) (*entity.User, string, string, error) {

	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", "", ErrInvalidCredential
	}

	if err := utils.CheckPassword(password, user.Password); err != nil {
		return nil, "", "", ErrInvalidCredential
	}

	accessToken, err := utils.GenerateToken(
		user.ID,
		s.jwtSecret,
		s.accessExpire,
	)
	if err != nil {
		return nil, "", "", NewAppError("TOKEN_FAILED", "failed generate token", 500, err)
	}

	refreshToken, err := utils.GenerateToken(
		user.ID,
		s.jwtSecret,
		s.refreshExpire,
	)
	if err != nil {
		return nil, "", "", NewAppError("TOKEN_FAILED", "failed generate refresh token", 500, err)
	}

	err = s.refreshTokenRepo.Save(ctx, &entity.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(time.Minute * time.Duration(s.refreshExpire)),
	})
	if err != nil {
		return nil, "", "", NewAppError("TOKEN_STORE_FAILED", "failed store refresh token", 500, err)
	}

	return user, accessToken, refreshToken, nil
}

func (s *authService) Register(ctx context.Context, name, email, password string) (*entity.User, error) {

	// 1. check email exists
	if _, err := s.userRepo.FindByEmail(ctx, email); err == nil {
		return nil, ErrEmailAlreadyUsed
	}

	// 2. hash password
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return nil, NewAppError(
			"PASSWORD_HASH_FAILED",
			"failed to process password",
			500,
			err,
		)
	}

	user := &entity.User{
		Name:     name,
		Email:    email,
		Password: hashed,
		IsActive: true,
	}

	// 3. transaction
	err = s.db.Transaction(func(tx *gorm.DB) error {
		if err := s.userRepo.CreatedAt(ctx, user); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, NewAppError(
			"REGISTER_FAILED",
			"failed to register user",
			500,
			err,
		)
	}

	return user, nil
}

func (s *authService) Refresh(ctx context.Context, refreshToken string) (string, error) {

	rt, err := s.refreshTokenRepo.FindByToken(ctx, refreshToken)
	if err != nil {
		return "", ErrUnauthorized
	}

	accessToken, err := utils.GenerateToken(
		rt.UserID,
		s.jwtSecret,
		s.accessExpire,
	)
	if err != nil {
		return "", NewAppError("TOKEN_FAILED", "failed generate token", 500, err)
	}

	return accessToken, nil
}

func (s *authService) Logout(ctx context.Context, userID int64) error {
	return s.refreshTokenRepo.DeleteByUserId(ctx, userID)
}
