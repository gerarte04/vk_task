package usecases

import (
	"context"
	"marketplace/internal/domain"
)

type AuthService interface {
	Login(ctx context.Context, login, password string) (string, error)
	Register(ctx context.Context, user *domain.User, password string) (*domain.User, error)
}
