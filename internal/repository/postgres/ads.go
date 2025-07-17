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

func (r *AdRepo) GetAdFeedWithOpts(ctx context.Context, opts domain.GetAdsOpts) ([]domain.FeedPageItem, error) {
	const op = "AdRepo.GetAdFeedWithOpts"

	query :=
		`SELECT *,
			CASE WHEN author_login = $1 THEN TRUE
				 ELSE FALSE
			END
			FROM ads WHERE price >= $2 AND price <= $3
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

	rows, err := r.pool.Query(ctx, query, opts.UserLogin, opts.LowerPrice, opts.HigherPrice)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	feed := make([]domain.FeedPageItem, 0, pageSize)

	for i := 1; rows.Next(); i++ {
		var ad domain.Ad
		item := domain.FeedPageItem{
			ItemNumber: i,
			Ad: &ad,
		}

		if err := rows.Scan(
			&ad.Id, &ad.AuthorLogin, &ad.Title, &ad.Description, &ad.ImageAddress, &ad.Price, &ad.CreationTime,
			&item.SelfAuthored,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		feed = append(feed, item)
	}

	return feed, nil
}
