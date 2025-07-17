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
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdRepo struct {
	pool *pgxpool.Pool
}

func NewAdRepo(pool *pgxpool.Pool) *AdRepo {
	return &AdRepo{
		pool: pool,
	}
}

func (r *AdRepo) PostAd(ctx context.Context, ad *domain.Ad) (uuid.UUID, error) {
	const op = "AdRepo.PostAd"

	query := 
		`INSERT INTO ads (author_login, title, description, image_address, price)
			VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var adId uuid.UUID
	err := r.pool.QueryRow(
		ctx, query, ad.AuthorLogin, ad.Title, ad.Description, ad.ImageAddress, ad.Price,
	).Scan(&adId)

	if err != nil {
		pgErr := pkgPostgres.DetectError(err)

		if errors.Is(pgErr, database.ErrForeignKeyViolation) {
			return uuid.Nil, fmt.Errorf("%s: %w", op, repository.ErrUserNotFound)
		}

		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return adId, nil
}

func (r *AdRepo) GetAdsWithOpts(ctx context.Context, opts domain.GetAdsOpts) ([]*domain.Ad, error) {
	const op = "AdRepo.GetAdsWithOpts"

	query :=
		`SELECT * FROM ads WHERE price >= $1 AND price <= $2
			ORDER BY %s %s
			LIMIT %d OFFSET %d`

	var orderColumn string
	switch opts.OrderOption {
	case domain.OrderByPrice:
		orderColumn = "price"
	default:
		orderColumn = "creation_time"
	}

	var orderDirection string
	switch opts.Ascending {
	case true:
		orderDirection = "ASC"
	default:
		orderDirection = "DESC"
	}

	const pageSize = 5 // need to config
	limit, offset := pageSize, pageSize * (opts.PageNumber - 1)

	query = fmt.Sprintf(query, orderColumn, orderDirection, limit, offset)

	rows, err := r.pool.Query(ctx, query, opts.LowerPrice, opts.HigherPrice)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	ads := make([]*domain.Ad, 0, pageSize)

	for rows.Next() {
		var ad domain.Ad

		if err := rows.Scan(
			&ad.Id, &ad.AuthorLogin, &ad.Title, &ad.Description, &ad.ImageAddress, &ad.Price, &ad.CreationTime,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		ads = append(ads, &ad)
	}

	return ads, nil
}
