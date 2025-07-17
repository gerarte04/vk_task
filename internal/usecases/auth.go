package usecases

import (
	"context"
	"marketplace/internal/domain"

	"github.com/google/uuid"
)

type AuthService interface {
	Login(ctx context.Context, login, password string) (string, error)
	Register(ctx context.Context, user *domain.User, password string) (uuid.UUID, error)
}
