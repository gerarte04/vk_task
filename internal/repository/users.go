package repository

import (
	"context"
	"marketplace/internal/domain"

	"github.com/google/uuid"
)

type UserRepo interface {
	PostUser(ctx context.Context, user *domain.User) (uuid.UUID, error)
	GetUserByLogin(ctx context.Context, login string) (*domain.User, error)
}
