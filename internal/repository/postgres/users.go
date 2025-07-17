package postgres

import (
	"context"
	"errors"
	"fmt"
	"marketplace/internal/domain"
	"marketplace/internal/repository"
	"marketplace/pkg/database"
	pkgPostgres "marketplace/pkg/database/postgres"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		pool: pool,
	}
}

func (r *UserRepo) PostUser(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	const op = "UserRepo.PostUser"

	query := "INSERT INTO users (login, password) VALUES ($1, $2) RETURNING id"

	var userId uuid.UUID
	err := r.pool.QueryRow(
		ctx, query, user.Login, user.PasswordHash,
	).Scan(&userId)

	if err != nil {
		pgErr := pkgPostgres.DetectError(err)

		if errors.Is(pgErr, database.ErrUniqueViolation) {
			return uuid.Nil, fmt.Errorf("%s: %w", op, repository.ErrUserExists)
		}

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return userId, nil
}

func (r *UserRepo) GetUserByLogin(ctx context.Context, login string) (*domain.User, error) {
	const op = "UserRepo.GetUserByLogin"

	query := "SELECT * FROM users WHERE login = $1"

	var user domain.User
	err := r.pool.QueryRow(
		ctx, query, login,
	).Scan(&user.Id, &user.Login, &user.PasswordHash)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrUserNotFound)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}
