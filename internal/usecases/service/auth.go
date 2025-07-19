package service

import (
	"context"
	"errors"
	"fmt"
	"marketplace/internal/config"
	"marketplace/internal/domain"
	"marketplace/internal/repository"
	"marketplace/internal/usecases"
	"time"

	pkgCrypto "marketplace/pkg/utils/crypto"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo repository.UserRepo
	jwtCfg config.JwtConfig
	cryptCfg config.CryptConfig
	publicKey []byte
	privateKey []byte
}

func NewAuthService(
	userRepo repository.UserRepo,
	jwtCfg config.JwtConfig,
	cryptCfg config.CryptConfig,
) (*AuthService, error) {
	const op = "NewAuthService"

	privateKey, err := pkgCrypto.ParsePrivateKeyFromPEM(jwtCfg.PrivateKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	publicKey, err := pkgCrypto.ParsePublicKeyFromPEM(jwtCfg.PublicKeyPEM)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &AuthService{
		userRepo: userRepo,
		jwtCfg: jwtCfg,
		cryptCfg: cryptCfg,
		privateKey: privateKey,
		publicKey: publicKey,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, login, password string) (string, error) {
	const op = "AuthService.Login"

	user, err := s.userRepo.GetUserByLogin(ctx, login)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", fmt.Errorf("%s: %w", op, usecases.ErrWrongPassword)
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	token, err := pkgCrypto.GenerateJwtToken(pkgCrypto.TokenClaims{
		Iss: s.jwtCfg.Issuer,
		Sub: user.Login,
		Exp: time.Now().Add(s.jwtCfg.ExpirationTime),
	}, s.privateKey)

	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return token, nil
}

func (s *AuthService) Register(ctx context.Context, user *domain.User, password string) (*domain.User, error) {
	const op = "AuthService.Register"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), s.cryptCfg.HashingCost)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	user.PasswordHash = passwordHash

	userId, err := s.userRepo.PostUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	user.Id = userId

	return user, nil
}
